package main

import (
	"errors"
	"math"
	"math/rand"
	"reflect"
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
	t          int
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
	// Read QAP problem
	n, A, B, err := ReadData(data)
	if err != nil {
		return nil, errors.New("error reading data (" + data + "): " + err.Error())
	}

	// Create population
	pop := NewPopulation(individuals, generations, n)

	return &evolutionaryAlgorithm{pop, 0, n, A, B}, nil
}

func NewPopulation(individuals, generations, solSize int) *Population {
	rand.Seed(time.Now().UnixNano())
	p := &Population{make([]*Individual, 0), generations, 0, nil}

	for i := 0; i < individuals; i++ {
		var ind *Individual = NewIndividual(solSize)
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

	return p
}

func NewIndividual(solSize int) *Individual {
	sols := make([]int, solSize)

	for i := range sols {
		sols[i] = -1
	}

	return &Individual{sols, 0, false}
}

func (ev *evolutionaryAlgorithm) Run(alg algorithmType) {
	ev.t = 0

	switch alg {
	case Generic:
		genericAlgorithm()
	case Baldwinian:
		baldwinianAlgorithm()
	case Lamarckian:
		lamarckianAlgorithm()
	}
}

func genericAlgorithm() {
}

func baldwinianAlgorithm() {
}

func lamarckianAlgorithm() {
}

func (ev *evolutionaryAlgorithm) PopulationSize() int {
	return len(ev.Population.Individuals)
}

// Fathers selection (generational)
func (ev *evolutionaryAlgorithm) SelectTournament() {
	rand.Seed(time.Now().UnixNano())
	p_selection := make([]*Individual, ev.PopulationSize())

	for i := range ev.Population.Individuals {
		father1 := rand.Intn(ev.PopulationSize() + 1)
		father2 := rand.Intn(ev.PopulationSize() + 1)

		if ev.Fitness(father1) > ev.Fitness(father2) {
			p_selection[i] = ev.Population.Individuals[father1]
		} else {
			p_selection[i] = ev.Population.Individuals[father2]
		}
	}

	ev.Population.Individuals = p_selection
}

func (ev *evolutionaryAlgorithm) OrderCrossover(crossPoint1, crossPoint2 int) {
	probCross := 0.8
	numIndividuals := int(math.Ceil(float64(ev.PopulationSize()) * probCross))
	p_cross := make([]*Individual, 0)
	bestFather := ev.Population.Individuals[0]

	//rand.Seed(time.Now().UnixNano())
	//rand.Intn(ev.n)
	//rand.Intn(ev.n-crossPoint1) + crossPoint1

	for i := 0; i < numIndividuals-1; i++ {
		son1 := NewIndividual(ev.n)
		son2 := NewIndividual(ev.n)
		father1 := ev.Population.Individuals[i+i]
		father2 := ev.Population.Individuals[i+i+1]
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

		// Elitism
		if ev.Fitness(i+i) < bestFather.Fitness {
			bestFather = father1
		} else if ev.Fitness(i+i+1) < bestFather.Fitness {
			bestFather = father2
		}

		son1.NeedFitness = true
		son2.NeedFitness = true
		p_cross = append(p_cross, son1, son2)
	}

	ev.Population.Individuals = p_cross
}

func (ev *evolutionaryAlgorithm) ExchangeMutation(point1, point2 int) {
	rand.Seed(time.Now().UnixNano())

	for _, ind := range ev.Population.Individuals {
		// 5% chance of mutation
		if rand.Float64() < 0.05 {
			ind.Solution[point1], ind.Solution[point2] = ind.Solution[point2], ind.Solution[point1]
			ind.NeedFitness = true
		}
	}
}

func (ev *evolutionaryAlgorithm) Elitism() {
	eliteExists := false
	worstFitness := -int(^uint(0) >> 1)
	i_worst := 0

	for i, v := range ev.Population.Individuals {
		if reflect.DeepEqual(ev.Population.BestFather.Solution, v.Solution) {
			eliteExists = true
			break
		}

		if ev.Fitness(i) > worstFitness {
			i_worst = i
		}
	}

	if !eliteExists {
		ev.Population.Individuals[i_worst] = ev.Population.BestFather
	}
}

func (ev *evolutionaryAlgorithm) Evaluate() {

}

func (ev *evolutionaryAlgorithm) Fitness(ind int) int {
	fitness := ev.Population.Individuals[ind].Fitness

	if ev.Population.Individuals[ind].NeedFitness {
		fitness = 0
		in := ev.Population.Individuals[ind]

		for i := 0; i < ev.n; i++ {
			for j := 0; j < ev.n; j++ {
				fitness += ev.A[i][j] * ev.B[in.Solution[i]][in.Solution[j]]
			}
		}

		ev.Population.Evaluations++
		ev.Population.Individuals[ind].NeedFitness = false
	}

	return fitness
}

func (ev *evolutionaryAlgorithm) BestSolution() ([]int, int) {
	solution := make([]int, ev.n)
	fitness := int(^uint(0) >> 1)

	for i := range ev.Population.Individuals {
		if ev.Fitness(i) < fitness {
			fitness = ev.Fitness(i)
			solution = ev.Population.Individuals[i].Solution
		}
	}

	return solution, fitness
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
