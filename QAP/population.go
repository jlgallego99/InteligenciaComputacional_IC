package main

type Population struct {
	Individuals []Individual
	Generations int
	BestFit     int
}

type Individual struct {
	Solution []int
}

func NewPopulation() *Population {
	return &Population{}
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

func (p *Population) OrderCrossover() {

}

func (p *Population) ExchangeMutation() {

}

func (p *Population) Evaluate() {

}

func (p *Population) Fitness(ind, n int, A, B [][]int) int {
	return p.Individuals[ind].Fitness(n, A, B)
}

func (in *Individual) Fitness(n int, A, B [][]int) int {
	fitness := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fitness += A[i][j] * B[in.Solution[i]][in.Solution[j]]
		}
	}

	return fitness
}

func (p *Population) BestFitness() []int {
	return nil
}

func (p *Population) BestSolution() []int {
	return nil
}
