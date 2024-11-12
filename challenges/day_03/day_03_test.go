package day_03

import (
	"github.com/collinbus/advent-of-code-2015/helpers"
	"testing"
)

func TestDay3Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	visitedHouses := Day3Part1(input)

	if visitedHouses != 2081 {
		t.Fatalf("%d is not the correct answer\n", visitedHouses)
	}
}

func TestDay3Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	visitedHouses := Day3Part2(input)

	if visitedHouses !=  2341 {
		t.Fatalf("%d is not the correct answer\n", visitedHouses)
	}
}

