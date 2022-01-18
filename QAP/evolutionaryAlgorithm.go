package main

type algorithmType int

const (
	Generico algorithmType = iota
	Baldwiniano
	Lamarckiano
)

type evolutionaryAlgorithm struct {
}

func (ev *evolutionaryAlgorithm) Run() {
	// TODO
}
