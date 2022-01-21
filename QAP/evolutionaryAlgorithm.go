package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
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

	// Crear vector con valores del 1 al n y hacerle shuffle

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
	// Initial population
	ev.Population = NewPopulation(ev.PopulationSize(), ev.Population.Generations, ev.n)

	if alg == Lamarckian || alg == Baldwinian {
		ev.twoOpt(alg)
	}

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

		if alg == Baldwinian || alg == Lamarckian {
			ev.twoOpt(alg)
		}

		ev.saveBestIndividual()

		_, fitness := ev.BestSolution()
		fmt.Println(t, fitness)

		WriteResults(t, fitness, f)
	}
	fmt.Println("")
}

func (ev *evolutionaryAlgorithm) saveBestIndividual() {
	copy(ev.Population.BestFather.Solution, ev.Population.Individuals[0].Solution)
	ev.Fitness(ev.Population.Individuals[0])
	ev.Population.BestFather.Fitness = ev.Population.Individuals[0].Fitness
	ev.Population.BestFather.NeedFitness = false

	for _, ind := range ev.Population.Individuals {
		if ev.Fitness(ind) < ev.Population.BestFather.Fitness {
			copy(ev.Population.BestFather.Solution, ind.Solution)
			ev.Population.BestFather.Fitness = ind.Fitness
			ev.Population.BestFather.NeedFitness = false
		}
	}
}

func (ev *evolutionaryAlgorithm) PopulationSize() int {
	return len(ev.Population.Individuals)
}

// Optimization for all individuals
func (ev *evolutionaryAlgorithm) twoOpt(alg algorithmType) {
	// Mirar hacer concurrente

	for _, S := range ev.Population.Individuals {
		S.NeedFitness = true
		ev.Fitness(S)
		optimized := false
		best := NewIndividual(ev.n)

		// Keep iterating n times or until the individual can't be optimized more
		for it := 0; it < ev.n && !optimized; it++ {
			copy(best.Solution, S.Solution)
			best.Fitness = S.Fitness
			best.NeedFitness = false

			for i := 0; i < ev.n; i++ {
				for j := i + 1; j < ev.n; j++ {
					T := NewIndividual(ev.n)
					copy(T.Solution, S.Solution)
					T.Fitness = S.Fitness
					T.NeedFitness = false

					T.Solution[i], T.Solution[j] = T.Solution[j], T.Solution[i]
					ev.RecalculateFitness(T, S, i, j)
					T.NeedFitness = false

					if ev.Fitness(T) < ev.Fitness(S) {
						// Lamarckian: inherit solution
						if alg == Lamarckian {
							copy(S.Solution, T.Solution)
						}

						S.Fitness = T.Fitness
						S.NeedFitness = false
					}
				}
			}

			optimized = ev.Fitness(best) > ev.Fitness(S)
		}
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

	copy(ev.Population.Individuals, p_cross)
}

func (ev *evolutionaryAlgorithm) ExchangeMutation() {
	for _, ind := range ev.Population.Individuals {
		point1 := rand.Intn(ev.n)
		point2 := rand.Intn(ev.n-point1) + point1

		// 25% chance of mutation
		if rand.Float64() < 0.25 {
			ind.Solution[point1], ind.Solution[point2] = ind.Solution[point2], ind.Solution[point1]
			ind.NeedFitness = true
		}
	}
}

func (ev *evolutionaryAlgorithm) Elitism() {
	eliteExists := false
	ev.Fitness(ev.Population.Individuals[0])
	worstFitness := ev.Population.Individuals[0].Fitness
	i_worst := 0

	for i, ind := range ev.Population.Individuals {
		if reflect.DeepEqual(ev.Population.BestFather.Solution, ind.Solution) {
			eliteExists = true
			break
		}

		if ev.Fitness(ind) > worstFitness {
			worstFitness = ind.Fitness
			i_worst = i
		}
	}

	if !eliteExists {
		copy(ev.Population.Individuals[i_worst].Solution, ev.Population.BestFather.Solution)
		ev.Population.Individuals[i_worst].Fitness = ev.Population.BestFather.Fitness
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

func (ev *evolutionaryAlgorithm) RecalculateFitness(ind, S *Individual, pos_a, pos_b int) {
	newFitness := ind.Fitness

	for i := 0; i < ev.n; i++ {
		newFitness -= ev.A[pos_a][i] * ev.B[S.Solution[pos_a]][S.Solution[i]]
		newFitness += ev.A[pos_a][i] * ev.B[ind.Solution[pos_a]][ind.Solution[i]]

		newFitness -= ev.A[pos_b][i] * ev.B[S.Solution[pos_b]][S.Solution[i]]
		newFitness += ev.A[pos_b][i] * ev.B[ind.Solution[pos_b]][ind.Solution[i]]

		if i != pos_a && i != pos_b {
			newFitness -= ev.A[i][pos_a] * ev.B[S.Solution[i]][S.Solution[pos_a]]
			newFitness += ev.A[i][pos_a] * ev.B[ind.Solution[i]][ind.Solution[pos_a]]

			newFitness -= ev.A[i][pos_b] * ev.B[S.Solution[i]][S.Solution[pos_b]]
			newFitness += ev.A[i][pos_b] * ev.B[ind.Solution[i]][ind.Solution[pos_b]]
		}
	}

	ind.Fitness = newFitness
	ind.NeedFitness = false
}

func (ev *evolutionaryAlgorithm) BestSolution() ([]int, int) {
	return ev.Population.BestFather.Solution, ev.Population.BestFather.Fitness
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
