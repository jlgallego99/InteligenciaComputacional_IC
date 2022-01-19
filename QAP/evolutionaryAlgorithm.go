package main

import (
	"errors"
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

func (ev *evolutionaryAlgorithm) Run(alg algorithmType) *Population {
	ev.t = 0

	switch alg {
	case Generic:
		return genericAlgorithm()
	case Baldwinian:
		return baldwinianAlgorithm()
	case Lamarckian:
		return lamarckianAlgorithm()
	}

	return nil
}

func genericAlgorithm() *Population {
	return nil
}

func baldwinianAlgorithm() *Population {
	return nil
}

func lamarckianAlgorithm() *Population {
	return nil
}
