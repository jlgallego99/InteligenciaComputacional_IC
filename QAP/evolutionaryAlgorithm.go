package main

type algorithmType int

const (
	Generic algorithmType = iota
	Baldwinian
	Lamarckian
)

type evolutionaryAlgorithm struct {
}

func NewEvolutionaryAlgorithm(n int, A, B [][]int) *evolutionaryAlgorithm {
	return &evolutionaryAlgorithm{}
}

func (ev *evolutionaryAlgorithm) Run(alg algorithmType) {
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
