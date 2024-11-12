package day_04

import (
	"testing"
)

func TestDay4Part1(t *testing.T) {
	input := "bgvyzdsv"

	visitedHouses := Day4Part(input, "00000")

	if visitedHouses !=  254575 {
		t.Fatalf("%d is not the correct answer\n", visitedHouses)
	}
}

func TestDay4Part1(t *testing.T) {
	input := "bgvyzdsv"

	visitedHouses := Day4Part(input, "000000")

	if visitedHouses != 1038736 {
		t.Fatalf("%d is not the correct answer\n", visitedHouses)
	}
}

