package day_05

import (
	"github.com/collinbus/advent-of-code-2015/helpers"
	"testing"
)

func TestDay05Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	result := Day05Part1(input)

	if result != 236 {
		t.Fatalf("%d is not the correct answer\n", result)
	}
}

func TestDay05Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	result := Day05Part2(input)

	if result != 0 {
		t.Fatalf("%d is not the correct position\n", result)
	}
}
