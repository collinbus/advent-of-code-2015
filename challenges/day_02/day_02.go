package day_02

import (
	"sort"
	"strings"
	"strconv"
)

func Day2Part1(lines []string) int {
	wrappingPaper := 0
	for _, line := range lines {
		dimensions := strings.Split(line, "x")
		length, _ := strconv.Atoi(dimensions[0])
		width, _ := strconv.Atoi(dimensions[1])
		height, _ := strconv.Atoi(dimensions[2])

		wrappingPaper += (2 * length * width)
		wrappingPaper += (2 * width * height)
		wrappingPaper += (2 * height * length)
		wrappingPaper += min(min(length * width, width * height), height * length)
	}
	return wrappingPaper
}

func Day2Part2(lines []string) int {
	ribbon := 0
	for _, line := range lines {
		stringDimensions := strings.Split(line, "x")
		first, _ := strconv.Atoi(stringDimensions[0])
		second, _ := strconv.Atoi(stringDimensions[1])
		third, _ := strconv.Atoi(stringDimensions[2])
		dimensions := []int{first, second, third}
		sort.Slice(dimensions, func(i, j int) bool { return dimensions[i] < dimensions[j] })

		ribbon += (2 * dimensions[0]) + (2 * dimensions[1])
		ribbon += (dimensions[0] * dimensions[1] * dimensions[2])
	}
	return ribbon
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

