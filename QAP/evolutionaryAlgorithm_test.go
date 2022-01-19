package main

import (
	"reflect"
	"testing"
)

var father1 []int = []int{1, 0, 4, 2, 6, 5, 3}
var father2 []int = []int{3, 1, 2, 0, 5, 6, 4}
var son1 []int = []int{0, 5, 4, 2, 6, 3, 1}
var son2 []int = []int{4, 6, 2, 0, 5, 3, 1}
var mut1 []int = []int{0, 5, 6, 2, 4, 3, 1}
var mut2 []int = []int{4, 6, 5, 0, 2, 3, 1}
var population *Population = &Population{[]*Individual{{father1}, {father2}}, 5, 0}
var ev = &evolutionaryAlgorithm{population, 0, 7, nil, nil}

func TestSelectTournament(t *testing.T) {

}

func TestOrderCrossover(t *testing.T) {
	ev.OrderCrossover(2, 4)

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
	ev.ExchangeMutation(2, 4)

	if ev.PopulationSize() != 2 {
		t.Errorf("Expected %v, got %v", 2, ev.PopulationSize())
	}

	if !reflect.DeepEqual(ev.Population.Individuals[0].Solution, mut1) {
		t.Errorf("Expected %v, got %v", son1, ev.Population.Individuals[0].Solution)
	}

	if !reflect.DeepEqual(ev.Population.Individuals[1].Solution, mut2) {
		t.Errorf("Expected %v, got %v", son2, ev.Population.Individuals[1].Solution)
	}
}
