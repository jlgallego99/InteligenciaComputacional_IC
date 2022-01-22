package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type algorithmType int

const (
	Generic algorithmType = iota
	Baldwinian
	Lamarckian
)

type evolutionaryAlgorithm struct {
	Population *Population
	n          int
	A          [][]int
	B          [][]int
}

type Population struct {
	Individuals []*Individual
	Generations int
	Evaluations int
	BestFather  *Individual
}

type Individual struct {
	Solution    []int
	Fitness     int
	NeedFitness bool
}

func NewEvolutionaryAlgorithm(data string, individuals, generations int) (*evolutionaryAlgorithm, error) {
	rand.Seed(time.Now().UnixNano())

	// Read QAP problem
	n, A, B, err := ReadData(data)
	if err != nil {
		return nil, errors.New("error reading data (" + data + "): " + err.Error())
	}

	// Create population
	pop := NewPopulation(individuals, generations, n)

	return &evolutionaryAlgorithm{pop, n, A, B}, nil
}

func NewPopulation(individuals, generations, solSize int) *Population {
	p := &Population{make([]*Individual, 0), generations, 0, nil}

	for i := 0; i < individuals; i++ {
		var ind *Individual = NewIndividual(solSize)
		ind.NeedFitness = true
		j := 0
		for j != solSize {
			val := rand.Intn(solSize)
			exist := false

			for _, v := range ind.Solution {
				if v == val {
					exist = true
					break
				}
			}

			if !exist {
				ind.Solution[j] = val
				j++
			}
		}

		p.Individuals = append(p.Individuals, ind)
	}

	p.BestFather = NewIndividual(solSize)

	return p
}

func NewIndividual(solSize int) *Individual {
	sols := make([]int, solSize)

	for i := range sols {
		sols[i] = -1
	}

	return &Individual{sols, 0, true}
}

func (ev *evolutionaryAlgorithm) Run(alg algorithmType) {
	// Start from optimized population
	ev.twoOptConcurrent(alg)

	// Results file
	name := ""
	switch alg {
	case Generic:
		name += "generic_" + strconv.Itoa(ev.PopulationSize()) + "_" + strconv.Itoa(ev.Population.Generations) + ".txt"

	case Baldwinian:
		name += "baldwinian_" + strconv.Itoa(ev.PopulationSize()) + "_" + strconv.Itoa(ev.Population.Generations) + ".txt"

	case Lamarckian:
		name += "lamarckian_" + strconv.Itoa(ev.PopulationSize()) + "_" + strconv.Itoa(ev.Population.Generations) + ".txt"

	}
	f := OpenResultsFile(name)
	defer f.Close()

	// Loop for generations
	fmt.Println("Generation: ")
	ev.saveBestIndividual()
	for t := 0; t < ev.Population.Generations; t++ {
		ev.SelectTournament()

		ev.OrderCrossover()

		ev.ExchangeMutation()

		ev.Elitism()

		ev.twoOptConcurrent(alg)

		ev.saveBestIndividual()

		_, fitness := ev.BestSolution()
		fmt.Println(t, fitness)

		WriteResults(t, fitness, f)
	}
	fmt.Println("")
}

func (ev *evolutionaryAlgorithm) twoOptConcurrent(alg algorithmType) {
	if alg == Lamarckian || alg == Baldwinian {
		var wg sync.WaitGroup

		// Launch threads
		numThreads := 10
		for i := 0; i < numThreads; i++ {
			tam := len(ev.Population.Individuals) / numThreads
			individuals := make([]*Individual, tam)
			copy(individuals, ev.Population.Individuals[i*tam:(i*tam)+tam])

			wg.Add(1)
			go ev.twoOpt(alg, individuals, &wg)
		}

		wg.Wait()
	}
}

// Optimization for all individuals
func (ev *evolutionaryAlgorithm) twoOpt(alg algorithmType, individuals []*Individual, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	for _, ind := range individuals {
		ev.Fitness(ind)
		S := NewIndividual(ev.n)
		copy(S.Solution, ind.Solution)
		S.Fitness = ev.Fitness(ind)
		S.NeedFitness = false

		notOptimized := true
		best := NewIndividual(ev.n)

		// Keep iterating n times or until the individual is optimized
		for it := 0; it < ev.n && notOptimized; it++ {
			copy(best.Solution, S.Solution)
			best.Fitness = ev.Fitness(S)
			best.NeedFitness = false

			for i := 0; i < ev.n; i++ {
				for j := i + 1; j < ev.n; j++ {
					T := NewIndividual(ev.n)
					copy(T.Solution, S.Solution)
					T.Fitness = ev.Fitness(S)
					T.NeedFitness = false

					T.Solution[i], T.Solution[j] = T.Solution[j], T.Solution[i]
					T.NeedFitness = true

					if ev.Fitness(T) < ev.Fitness(S) {
						copy(S.Solution, T.Solution)
						S.Fitness = ev.Fitness(T)
						S.NeedFitness = false
					}
				}
			}

			notOptimized = ev.Fitness(best) > ev.Fitness(S)
		}

		// Lamarckian: inherit solution
		if alg == Lamarckian {
			copy(ind.Solution, S.Solution)
		}
		ind.Fitness = ev.Fitness(S)
		ind.NeedFitness = false
	}
}

// Fathers selection (generational)
func (ev *evolutionaryAlgorithm) SelectTournament() {
	p_selection := make([]*Individual, 0)

	for range ev.Population.Individuals {
		father1 := rand.Intn(ev.PopulationSize())
		father2 := rand.Intn(ev.PopulationSize())
		for father2 != father1 {
			father2 = rand.Intn(ev.PopulationSize())
		}

		if ev.Fitness(ev.Population.Individuals[father1]) < ev.Fitness(ev.Population.Individuals[father2]) {
			newFather := NewIndividual(ev.n)
			copy(newFather.Solution, ev.Population.Individuals[father1].Solution)
			newFather.NeedFitness = true
			p_selection = append(p_selection, newFather)
		} else {
			newFather := NewIndividual(ev.n)
			copy(newFather.Solution, ev.Population.Individuals[father2].Solution)
			newFather.NeedFitness = true
			p_selection = append(p_selection, newFather)
		}
	}

	copy(ev.Population.Individuals, p_selection)
}

func (ev *evolutionaryAlgorithm) OrderCrossover() {
	probCross := 0.8
	numIndividuals := int(math.Ceil(float64(ev.PopulationSize()) * probCross))
	p_cross := make([]*Individual, 0)
	bestFather := NewIndividual(ev.n)
	copy(bestFather.Solution, ev.Population.Individuals[0].Solution)
	bestFather.NeedFitness = true

	for i := 0; i < numIndividuals/2; i++ {
		son1 := NewIndividual(ev.n)
		son2 := NewIndividual(ev.n)
		father1 := NewIndividual(ev.n)
		father2 := NewIndividual(ev.n)
		copy(father1.Solution, ev.Population.Individuals[i+i].Solution)
		copy(father2.Solution, ev.Population.Individuals[i+i+1].Solution)
		father1.NeedFitness = true
		father2.NeedFitness = true
		crossPoint1 := rand.Intn(ev.n)
		crossPoint2 := rand.Intn(ev.n-crossPoint1) + crossPoint1

		copy(son1.Solution[crossPoint1:crossPoint2+1%ev.n], father1.Solution[crossPoint1:crossPoint2+1%ev.n])
		copy(son2.Solution[crossPoint1:crossPoint2+1%ev.n], father2.Solution[crossPoint1:crossPoint2+1%ev.n])

		index := (crossPoint2 + 1) % ev.n
		for j := 0; j < ev.n || index < crossPoint1; j++ {
			fi := (crossPoint2 + 1 + j) % ev.n

			if !contains(son1.Solution, father2.Solution[fi]) {
				son1.Solution[index] = father2.Solution[fi]
				index = (index + 1) % ev.n
			}
		}

		index = (crossPoint2 + 1) % ev.n
		for j := 0; j < ev.n || index < crossPoint1; j++ {
			fi := (crossPoint2 + 1 + j) % ev.n

			if !contains(son2.Solution, father1.Solution[fi]) {
				son2.Solution[index] = father1.Solution[fi]
				index = (index + 1) % ev.n
			}
		}

		son1.NeedFitness = true
		son2.NeedFitness = true
		p_cross = append(p_cross, son1, son2)
	}

	for i := numIndividuals; i < ev.PopulationSize(); i++ {
		p_cross = append(p_cross, ev.Population.Individuals[i])
	}

	copy(ev.Population.Individuals, p_cross)
}

func (ev *evolutionaryAlgorithm) ExchangeMutation() {
	for _, ind := range ev.Population.Individuals {
		point1 := rand.Intn(ev.n)
		point2 := rand.Intn(ev.n-point1) + point1

		// 50% chance of mutation
		if rand.Float64() < 0.5 {
			ind.Solution[point1], ind.Solution[point2] = ind.Solution[point2], ind.Solution[point1]
			ind.NeedFitness = true
		}
	}
}

func (ev *evolutionaryAlgorithm) saveBestIndividual() {
	copy(ev.Population.BestFather.Solution, ev.Population.Individuals[0].Solution)
	ev.Fitness(ev.Population.Individuals[0])
	ev.Population.BestFather.Fitness = ev.Fitness(ev.Population.Individuals[0])
	ev.Population.BestFather.NeedFitness = false

	for _, ind := range ev.Population.Individuals {
		if ev.Fitness(ind) < ev.Fitness(ev.Population.BestFather) {
			copy(ev.Population.BestFather.Solution, ind.Solution)
			ev.Population.BestFather.Fitness = ev.Fitness(ind)
			ev.Population.BestFather.NeedFitness = false
		}
	}
}

func (ev *evolutionaryAlgorithm) Elitism() {
	eliteExists := false
	ev.Fitness(ev.Population.Individuals[0])
	worstFitness := ev.Fitness(ev.Population.Individuals[0])
	i_worst := 0

	for i, ind := range ev.Population.Individuals {
		if reflect.DeepEqual(ev.Population.BestFather.Solution, ind.Solution) {
			eliteExists = true
			break
		}

		if ev.Fitness(ind) > worstFitness {
			worstFitness = ev.Fitness(ind)
			i_worst = i
		}
	}

	if !eliteExists {
		copy(ev.Population.Individuals[i_worst].Solution, ev.Population.BestFather.Solution)
		ev.Population.Individuals[i_worst].Fitness = ev.Fitness(ev.Population.BestFather)
		ev.Population.Individuals[i_worst].NeedFitness = false
	}
}

func (ev *evolutionaryAlgorithm) Fitness(ind *Individual) int {
	fitness := ind.Fitness

	if ind.NeedFitness {
		fitness = 0

		for i := 0; i < ev.n; i++ {
			for j := 0; j < ev.n; j++ {
				fitness = fitness + (ev.A[i][j] * ev.B[ind.Solution[i]][ind.Solution[j]])
			}
		}

		ev.Population.Evaluations++
		ind.Fitness = fitness
		ind.NeedFitness = false
	}

	return fitness
}

func (ev *evolutionaryAlgorithm) BestSolution() ([]int, int) {
	return ev.Population.BestFather.Solution, ev.Fitness(ev.Population.BestFather)
}

func (ev *evolutionaryAlgorithm) PopulationSize() int {
	return len(ev.Population.Individuals)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
