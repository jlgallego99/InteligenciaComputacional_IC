package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const bestKnownFitness int = 44759294

func fitnessDifference(c int) float32 {
	difference := float32(5.0 - 100.0*(c-bestKnownFitness)/bestKnownFitness)

	if difference < 0 {
		difference = 0.0
	}

	return difference
}

func main() {
	data := "chr12a.dat"
	individuals, _ := strconv.Atoi(os.Args[1])
	generations, _ := strconv.Atoi(os.Args[2])

	// Run evolutionary algorithm
	ev, err := NewEvolutionaryAlgorithm(data, individuals, generations)
	if err != nil {
		_ = fmt.Errorf(err.Error())

		return
	}

	fmt.Println("Running evolutionary algorithms...")
	fmt.Println("Population size:", ev.PopulationSize())
	fmt.Println("Generations:", ev.Population.Generations)
	fmt.Println("")

	// Generic algorithm solution
	start := time.Now()
	ev.Run(Generic)
	end := time.Since(start)
	fmt.Println("GENERIC ALGORITHM")
	fmt.Println("Time (seconds):", end.Seconds())
	fmt.Println("Fitness:", ev.BestFitness())
	fmt.Println("Fitness difference from best known:", fitnessDifference(ev.BestFitness()))
	fmt.Println("Solution:", ev.BestSolution())

	// Baldwinian algorithm solution
	/*start = time.Now()
	ev.Run(Baldwinian)
	end = time.Since(start)
	fmt.Println("GENERIC ALGORITHM")
	fmt.Println("Time (seconds):", end.Seconds())
	fmt.Println("Fitness:", ev.BestFitness())
	fmt.Println("Fitness difference from best known:", fitnessDifference(ev.BestFitness()))
	fmt.Println("Solution:", ev.BestSolution())

	// Lamarckian algorithm solution
	start = time.Now()
	ev.Run(Lamarckian)
	end = time.Since(start)
	fmt.Println("GENERIC ALGORITHM")
	fmt.Println("Time (seconds):", end.Seconds())
	fmt.Println("Fitness:", ev.BestFitness())
	fmt.Println("Fitness difference from best known:", fitnessDifference(ev.BestFitness()))
	fmt.Println("Solution:", ev.BestSolution())*/
}
