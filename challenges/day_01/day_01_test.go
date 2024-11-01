package day_01

import (
	"github.com/collinbus/advent-of-code-2015/helpers"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	floor := Day1Part1(input)

	if floor != 232 {
		t.Fatalf("%d is not the correct answer\n", floor)
	}
}

func TestDay1Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	firstPosition := Day1Part2(input)

	if firstPosition != 1783 {
		t.Fatalf("%d is not the correct position\n", firstPosition)
	}
}

