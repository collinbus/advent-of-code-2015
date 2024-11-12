package day_03

import (
	"fmt"
)

func Day3Part1(lines []string) int {
	x, y := 0, 0
	visited := make(map[string]bool)
	visited["00"] = true
	for _, char := range lines[0] {
		switch char {
		case '<':
			x--
		case '>':
			x++
		case '^':
			y--
		default:
			y++
		}

		coordinates := fmt.Sprintf("%d%d", y, x)
		visited[coordinates] = true
	}
	return len(visited)
}

func Day3Part2(lines []string) int {
	x, y := 0, 0
	rx, ry := 0, 0
	visited := make(map[string]bool)
	visited["0,0"] = true
	for i, char := range lines[0] {
		switch char {
		case '<':
			if i % 2 == 0 {
				x--
			} else {
				rx--
			}
		case '>':
			if i % 2 == 0 {
				x++
			} else {
				rx++
			}
		case '^':
			if i % 2 == 0 {
				y++
			} else {
				ry++
			}
		default:
			if i % 2 == 0 {
				y--
			} else {
				ry--
			}
		}

		coordinates := fmt.Sprintf("%d,%d", y, x)
		rCoordinates := fmt.Sprintf("%d,%d", ry, rx)
		visited[coordinates] = true
		visited[rCoordinates] = true
	}
	return len(visited)
}

