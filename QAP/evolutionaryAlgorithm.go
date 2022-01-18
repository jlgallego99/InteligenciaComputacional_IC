package main

type algorithmType int

const (
	Generic algorithmType = iota
	Baldwinian
	Lamarckian
)

const bestKnownSolution int = 44759294

type evolutionaryAlgorithm struct {
	t int
}

func NewEvolutionaryAlgorithm(n int, A, B [][]int) *evolutionaryAlgorithm {
	return &evolutionaryAlgorithm{0}
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
