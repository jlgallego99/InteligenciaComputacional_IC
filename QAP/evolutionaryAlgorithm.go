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

const bestKnownSolution int = 44759294

type evolutionaryAlgorithm struct {
	t int
	n int
	A [][]int
	B [][]int
}

func NewEvolutionaryAlgorithm(data string) (*evolutionaryAlgorithm, error) {
	// Read QAP problem
	n, A, B, err := ReadData(data)

	if err != nil {
		return nil, errors.New("error reading data (" + data + "): " + err.Error())
	}

	return &evolutionaryAlgorithm{0, n, A, B}, nil
}

func (ev *evolutionaryAlgorithm) Run(alg algorithmType) *Population {
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
