package main

import (
	"errors"
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

func (ev *evolutionaryAlgorithm) OrderCrossover() {

}

func (ev *evolutionaryAlgorithm) ExchangeMutation() {

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
