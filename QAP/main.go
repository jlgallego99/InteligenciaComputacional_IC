package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const bestKnownFitness int = 44759294

func fitnessDifference(c int) float32 {
	difference := 5.0 - 100.0*(float32(c)-float32(bestKnownFitness))/float32(bestKnownFitness)

	if difference < 0 {
		difference = 0.0
	}

	return difference
}

func main() {
	data := "tai256c.dat"
	var alg, individuals, generations int

	if len(os.Args) != 4 {
		fmt.Println("Incorrect arguments...")
		fmt.Println("Default algorithm: generic")
		fmt.Println("Default individuals: 100")
		fmt.Println("Default generations: 100")
		fmt.Println("Usage for the next time:", os.Args[0], "type individuals generations")
		fmt.Println("Types: 0 (generic), 1 (baldwinian), 2 (lamarckian)")
		fmt.Println("")

		individuals = 100
		generations = 100
	} else {
		alg, _ = strconv.Atoi(os.Args[1])
		individuals, _ = strconv.Atoi(os.Args[2])
		generations, _ = strconv.Atoi(os.Args[3])
	}

	// Run evolutionary algorithm
	ev, err := NewEvolutionaryAlgorithm(data, individuals, generations)
	if err != nil {
		_ = fmt.Errorf(err.Error())

		return
	}

	fmt.Println("Running evolutionary algorithm...")
	fmt.Println("Population size:", ev.PopulationSize())
	fmt.Println("Generations:", ev.Population.Generations)
	fmt.Println("")

	switch algorithmType(alg) {
	case Generic:
		fmt.Println("GENERIC EVOLUTIONARY ALGORITHM")

		// Generic algorithm solution
		start := time.Now()
		ev.Run(Generic)
		end := time.Since(start)
		solution, fitness := ev.BestSolution()
		fmt.Println("GENERIC ALGORITHM")
		fmt.Println("Time (seconds):", end.Seconds())
		fmt.Println("Solution:", solution)
		fmt.Println("Fitness:", fitness)
		fmt.Println("Score from best known solution:", fitnessDifference(fitness))

	case Baldwinian:
		fmt.Println("BALDWINIAN EVOLUTIONARY ALGORITHM")

		// Baldwinian algorithm solution
		start := time.Now()
		ev.Run(Baldwinian)
		end := time.Since(start)
		solution, fitness := ev.BestSolution()
		fmt.Println("BALDWINIAN ALGORITHM")
		fmt.Println("Time (seconds):", end.Seconds())
		fmt.Println("Solution:", solution)
		fmt.Println("Fitness:", fitness)
		fmt.Println("Fitness difference from best known:", fitnessDifference(fitness))

	case Lamarckian:
		fmt.Println("LAMARCKIAN EVOLUTIONARY ALGORITHM")

		// Lamarckian algorithm solution
		start := time.Now()
		ev.Run(Lamarckian)
		end := time.Since(start)
		fmt.Println("LAMARCKIAN ALGORITHM")
		solution, fitness := ev.BestSolution()
		fmt.Println("Time (seconds):", end.Seconds())
		fmt.Println("Solution:", solution)
		fmt.Println("Fitness:", fitness)
		fmt.Println("Fitness difference from best known:", fitnessDifference(fitness))
	}
}
