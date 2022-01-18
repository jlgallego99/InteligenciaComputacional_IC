package main

import "fmt"

func main() {
	data := "chr12a.dat"

	// Read QAP problem
	n, A, B, err := ReadData(data)

	if err != nil {
		_ = fmt.Errorf("error reading data (%s): %v", data, err)

		return
	}

	// Run evolutionary algorithm
	ev := NewEvolutionaryAlgorithm(n, A, B)
	ev.Run(Generic)
	ev.Run(Baldwinian)
	ev.Run(Lamarckian)
}
