#!/bin/bash

day="day_$1"
folderName="challenges/$day"
mkdir $folderName


cat > "$folderName/${day}.go" << EOF
package $day

import (
	"fmt"
)

func Day${1}Part1(lines []string) int {
	fmt.Println("Implement me!")
	return -1
}

func Day${1}Part2(lines []string) int {
	fmt.Println("Implement me!")
	return -1
}
EOF

cat > "$folderName/${day}_test.go" << EOF
package $day

import (
	"github.com/collinbus/advent-of-code-2015/helpers"
	"testing"
)

func TestDay${1}Part1(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	result := Day${1}Part1(input)

	if result != 0 {
		t.Fatalf("%d is not the correct answer\n", result)
	}
}

/*func TestDay${1}Part2(t *testing.T) {
	input := helpers.ReadInput("input.txt")

	result := Day${1}Part2(input)

	if result != 0 {
		t.Fatalf("%d is not the correct answer\n", result)
	}
}*/
EOF
