package day_06

import (
	"github.com/collinbus/advent-of-code-2015/helpers"
	"testing"
)

func TestDay06Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	result := Day06Part1(input)

	if result != 543903 {
		t.Fatalf("%d is not the correct answer\n", result)
	}
}

func TestDay06Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	result := Day06Part2(input)

	if result != 0 {
		t.Fatalf("%d is not the correct answer\n", result)
	}
}
