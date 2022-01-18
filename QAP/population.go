package main

type Population struct {
	Individuals [][]int
	Generations int
	BestFit     int
}

func (p *Population) Size() int {
	return len(p.Individuals)
}

// Fathers selection
func (p *Population) SelectTournament() {

}

// Survivors selection
func (p *Population) Elitism() {

}

func (p *Population) OnePointCrossover() {

}

func (p *Population) UniformCrossover() {

}

func (p *Population) Mutate() {

}

func (p *Population) Evaluate() {

}

func (p *Population) Fitness(ind int) float64 {
	return 0
}

func (p *Population) BestFitness() []int {
	return nil
}
