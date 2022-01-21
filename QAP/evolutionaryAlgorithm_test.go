package main

import (
	"math/rand"
	"reflect"
	"testing"
)

var father1 []int = []int{1, 0, 4, 2, 6, 5, 3}
var father2 []int = []int{3, 1, 2, 0, 5, 6, 4}
var son1 []int = []int{1, 0, 4, 2, 5, 6, 3}
var son2 []int = []int{1, 4, 2, 0, 6, 5, 3}
var mut1 []int = []int{0, 5, 6, 2, 4, 3, 1}
var mut2 []int = []int{4, 6, 5, 0, 2, 3, 1}
var population *Population = &Population{[]*Individual{{father1, 0, true}, {father2, 0, true}}, 5, 0, nil}
var A [][]int = [][]int{
	{0, 1, 1, 1, 1, 1, 1},
	{1, 0, 1, 1, 1, 1, 1},
	{1, 1, 0, 1, 1, 1, 1},
	{1, 1, 1, 0, 1, 1, 1},
	{1, 1, 1, 1, 0, 1, 1},
	{1, 1, 1, 1, 1, 0, 1},
	{1, 1, 1, 1, 1, 1, 0},
}
var B [][]int = [][]int{
	{0, 1, 1, 1, 1, 1, 1},
	{1, 0, 1, 1, 1, 1, 1},
	{1, 1, 0, 1, 1, 1, 1},
	{1, 1, 1, 0, 1, 1, 1},
	{1, 1, 1, 1, 0, 1, 1},
	{1, 1, 1, 1, 1, 0, 1},
	{1, 1, 1, 1, 1, 1, 0},
}
var ev = &evolutionaryAlgorithm{population, 7, A, B}

func TestNewPopulation(t *testing.T) {
	pop := NewPopulation(5, 5, 5)

	if len(pop.Individuals) != 5 {
		t.Errorf("Expected %v, got %v", 5, len(pop.Individuals))
	}

	for _, ind := range pop.Individuals {
		if len(ind.Solution) != 5 {
			t.Errorf("Expected %v, got %v", 5, len(ind.Solution))
		}
	}
}

func TestOrderCrossover(t *testing.T) {
	rand.Seed(5)
	ev.OrderCrossover()

	if ev.PopulationSize() != 2 {
		t.Errorf("Expected %v, got %v", 2, ev.PopulationSize())
	}

	if !reflect.DeepEqual(ev.Population.Individuals[0].Solution, son1) {
		t.Errorf("Expected %v, got %v", son1, ev.Population.Individuals[0].Solution)
	}

	if !reflect.DeepEqual(ev.Population.Individuals[1].Solution, son2) {
		t.Errorf("Expected %v, got %v", son2, ev.Population.Individuals[1].Solution)
	}
}

func TestExchangeMutation(t *testing.T) {
	rand.Seed(1)
	ev.ExchangeMutation()

	if ev.PopulationSize() != 2 {
		t.Errorf("Expected %v, got %v", 2, ev.PopulationSize())
	}

	if !reflect.DeepEqual(ev.Population.Individuals[0].Solution, mut1) && !ev.Population.Individuals[0].NeedFitness {
		t.Errorf("Expected %v, got %v", mut1, ev.Population.Individuals[0].Solution)
	}

	if !reflect.DeepEqual(ev.Population.Individuals[1].Solution, mut2) && !ev.Population.Individuals[1].NeedFitness {
		t.Errorf("Expected %v, got %v", mut2, ev.Population.Individuals[1].Solution)
	}
}

var population2 = NewPopulation(3, 1, 3)
var A2 = [][]int{{0, 10, 30}, {10, 0, 20}, {30, 20, 0}}
var B2 = [][]int{{0, 1, 2}, {1, 0, 1}, {2, 1, 0}}
var ev2 = &evolutionaryAlgorithm{population2, 3, A2, B2}

func TestSelectTournament(t *testing.T) {
	ev2.SelectTournament()

	if ev2.PopulationSize() != 3 {
		t.Errorf("Expected %v, got %v", 3, ev2.PopulationSize())
	}
}

func TestTwoOpt(t *testing.T) {
	oldfit1 := ev2.Fitness(ev2.Population.Individuals[0])
	oldfit2 := ev2.Fitness(ev2.Population.Individuals[1])
	oldfit3 := ev2.Fitness(ev2.Population.Individuals[2])

	ev2.twoOpt()

	if !(ev2.Fitness(ev2.Population.Individuals[0]) <= oldfit1) {
		t.Errorf("Fitness not improved: %v < %v", ev2.Fitness(ev2.Population.Individuals[0]), oldfit1)
	}

	if !(ev2.Fitness(ev2.Population.Individuals[1]) <= oldfit2) {
		t.Errorf("Fitness not improved: %v < %v", ev2.Fitness(ev2.Population.Individuals[1]), oldfit2)
	}

	if !(ev2.Fitness(ev2.Population.Individuals[2]) <= oldfit3) {
		t.Errorf("Fitness not improved: %v < %v", ev2.Fitness(ev2.Population.Individuals[2]), oldfit3)
	}
}
