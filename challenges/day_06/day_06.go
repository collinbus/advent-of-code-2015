package day_06

import (
	"strconv"
	"strings"
)

type Instruction int

const (
	TURN_ON = iota
	TURN_OFF
	TOGGLE
)

func Day06Part1(lines []string) int {
	lightsLid := 0
	grid := [1000][1000]bool{}
	for _, line := range lines {
		startIndex := 0
		instruction := TURN_OFF
		if strings.HasPrefix(line,"turn on") {
			startIndex = 8
			instruction = TURN_ON
		} else if strings.HasPrefix(line, "toggle") {
			startIndex = 7
			instruction = TOGGLE
		} else {
			startIndex = 9
		}
		coordinates := strings.Split(line[startIndex:], " through ")
		startCoordinates := strings.Split(coordinates[0], ",")
		endCoordinates := strings.Split(coordinates[1], ",")
		sx, _ := strconv.Atoi(startCoordinates[0])	
		sy, _ := strconv.Atoi(startCoordinates[1])
		ex, _ := strconv.Atoi(endCoordinates[0])
		ey, _ := strconv.Atoi(endCoordinates[1])
		
		for y := sy; y <= ey ; y++ {
			for x := sx ; x <= ex ; x++ {
				if instruction == TURN_ON {
					grid[y][x] = true
				} else if instruction == TURN_OFF {
					grid[y][x] = false
				} else {
					grid[y][x] = !grid[y][x]
				}
			}
		}
	}

	for y := 0; y < 1000 ; y++ {
		for x := 0 ; x < 1000 ; x++ {
			if grid[y][x] {
				lightsLid++
			}
		}
	}	

	return lightsLid
}

func Day06Part2(lines []string) int {
	brightness := 0
	grid := [1000][1000]int{}
	for _, line := range lines {
		startIndex := 0
		instruction := TURN_OFF
		if strings.HasPrefix(line,"turn on") {
			startIndex = 8
			instruction = TURN_ON
		} else if strings.HasPrefix(line, "toggle") {
			startIndex = 7
			instruction = TOGGLE
		} else {
			startIndex = 9
		}
		coordinates := strings.Split(line[startIndex:], " through ")
		startCoordinates := strings.Split(coordinates[0], ",")
		endCoordinates := strings.Split(coordinates[1], ",")
		sx, _ := strconv.Atoi(startCoordinates[0])	
		sy, _ := strconv.Atoi(startCoordinates[1])
		ex, _ := strconv.Atoi(endCoordinates[0])
		ey, _ := strconv.Atoi(endCoordinates[1])
		
		for y := sy; y <= ey ; y++ {
			for x := sx ; x <= ex ; x++ {
				if instruction == TURN_ON {
					grid[y][x]++
				} else if instruction == TURN_OFF {
					newValue := grid[y][x] - 1
					grid[y][x] = max(newValue, 0)
				} else {
					grid[y][x]+= 2
				}
			}
		}
	}

	for y := 0; y < 1000 ; y++ {
		for x := 0 ; x < 1000 ; x++ {
			brightness += grid[y][x]
		}
	}	

	return brightness
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
