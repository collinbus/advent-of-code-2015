package day_02

import (
	"github.com/collinbus/advent-of-code-2015/helpers"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	wrappingPaper := Day2Part1(input)

	if wrappingPaper != 1588178 {
		t.Fatalf("%d is not the correct answer\n", wrappingPaper)
	}
}

func TestDay2Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	ribbon := Day2Part2(input)

	if ribbon != 3783758 {
		t.Fatalf("%d is not the correct answer\n", ribbon)
	}
}

