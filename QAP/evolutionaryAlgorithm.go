package main

import (
	"errors"
	"math"
	"math/rand"
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
	BestFit     int
}

type Individual struct {
	Solution []int
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
	p := &Population{make([]*Individual, 0), generations, 0}

	for i := 0; i < individuals; i++ {
		p.Individuals = append(p.Individuals, NewIndividual(solSize))
	}

	return p
}

func NewIndividual(solSize int) *Individual {
	sols := make([]int, solSize)

	for i := range sols {
		sols[i] = -1
	}

	return &Individual{sols}
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

// Survivors selection
func (ev *evolutionaryAlgorithm) Elitism() {

}

func (ev *evolutionaryAlgorithm) OrderCrossover(crossPoint1, crossPoint2 int) {
	probCross := 0.8
	numIndividuals := int(math.Ceil(float64(ev.PopulationSize()) * probCross))
	p_cross := make([]*Individual, 0)

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

		p_cross = append(p_cross, son1, son2)
	}

	ev.Population.Individuals = p_cross
}

func (ev *evolutionaryAlgorithm) ExchangeMutation(point1, point2 int) {
	for _, ind := range ev.Population.Individuals {
		ind.Solution[point1], ind.Solution[point2] = ind.Solution[point2], ind.Solution[point1]
	}
}

func (ev *evolutionaryAlgorithm) Evaluate() {

}

func (ev *evolutionaryAlgorithm) Fitness(ind int) int {
	fitness := 0
	in := ev.Population.Individuals[ind]

	for i := 0; i < ev.n; i++ {
		for j := 0; j < ev.n; j++ {
			fitness += ev.A[i][j] * ev.B[in.Solution[i]][in.Solution[j]]
		}
	}

	return fitness
}

func (ev *evolutionaryAlgorithm) BestFitness() int {
	return ev.Population.BestFit
}

func (ev *evolutionaryAlgorithm) BestSolution() []int {
	return nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
