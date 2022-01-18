package main

import (
	"fmt"
	"time"
)

func main() {
	data := "chr12a.dat"

	// Run evolutionary algorithm
	ev, err := NewEvolutionaryAlgorithm(data)
	if err != nil {
		_ = fmt.Errorf(err.Error())

		return
	}

	// Generic algorithm solution
	start := time.Now()
	popGeneric := ev.Run(Generic)
	end := time.Since(start)
	fmt.Println("GENERIC ALGORITHM")
	fmt.Println("Time (seconds): ", end.Seconds())
	fmt.Println("Fitness: ", popGeneric.BestFitness())
	fmt.Println("Solution: ", popGeneric.BestSolution())

	// Baldwinian algorithm solution
	start = time.Now()
	popBaldwinian := ev.Run(Baldwinian)
	end = time.Since(start)
	fmt.Println("GENERIC ALGORITHM")
	fmt.Println("Time (seconds): ", end.Seconds())
	fmt.Println("Fitness: ", popBaldwinian.BestFitness())
	fmt.Println("Solution: ", popBaldwinian.BestSolution())

	// Generic algorithm solution
	start = time.Now()
	popLamarckian := ev.Run(Lamarckian)
	end = time.Since(start)
	fmt.Println("GENERIC ALGORITHM")
	fmt.Println("Time (seconds): ", end.Seconds())
	fmt.Println("Fitness: ", popLamarckian.BestFitness())
	fmt.Println("Solution: ", popLamarckian.BestSolution())
}
