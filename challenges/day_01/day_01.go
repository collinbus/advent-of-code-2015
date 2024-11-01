package day_01

func Day1Part1(lines []string) int {
	floor := 0
	for _, char := range lines[0] {
		if char == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func Day1Part2(lines []string) int {
	floor := 0
	for i, char := range lines[0] {
		if char == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}

