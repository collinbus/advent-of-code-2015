package day_05

import (
	"fmt"
	"regexp"
)

func Day05Part1(lines []string) int {
	niceStrings := 0
	for _, line := range lines {
		if isNice(line) {
			niceStrings++
		}
	}
	return niceStrings
}

func Day05Part2(lines []string) int {
	fmt.Println("Implement me!")
	return -1
}

func isNice(s string) bool {
	// Check for vowels
	vowels := regexp.MustCompile("[aeiou]").FindAllString(s, -1)
	if len(vowels) < 3 {
		return false
	}

	// Check for disallowed substrings
	for _, bad := range []string{"ab", "cd", "pq", "xy"} {
		if regexp.MustCompile(bad).MatchString(s) {
			return false
		}
	}

	for i := 1 ; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}

	return false
}

