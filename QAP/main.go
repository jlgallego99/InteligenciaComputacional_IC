package main

import "fmt"

func main() {
	data := "chr12a.dat"

	// Run evolutionary algorithm
	ev, err := NewEvolutionaryAlgorithm(data)
	if err != nil {
		_ = fmt.Errorf(err.Error())

		return
	}

	ev.Run(Generic)
	ev.Run(Baldwinian)
	ev.Run(Lamarckian)
}
