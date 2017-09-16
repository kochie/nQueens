package main

import (
	"testing"
)

func TestGetFitnessDiagonal(t *testing.T) {
	board := []int{0, 1, 2, 3, 4}

	fitness := getFitness(&board)

	if fitness != 1 {
		t.Error("Fitness expected to be 1, got", fitness)
	}

	board = []int{0, 1, 1, 0, 0}

	fitness = getFitness(&board)
	if fitness != 1 {
		t.Error("Fitness expected to be 1, got", fitness)
	}
}

func TestGetFitnessPerfectFitness(t *testing.T) {
	board := []int{0,2,4,1,3}

	fitness := getFitness(&board)

	if fitness != 5 {
		t.Error("Fitness expected to be 5, got", fitness)
	}
}
