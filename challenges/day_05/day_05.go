package day_05

import (
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
	niceStrings := 0
	for _, line := range lines {
		if isNicePart2(line) {
			niceStrings++
		}
	}
	return niceStrings
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

func isNicePart2(s string) bool {
	hasTwoPairs := false
	for i := 0; i < len(s) - 1; i++ {
		for j := i + 2; j < len(s) - 1; j++ {
			if s[i:i+2] == s[j:j+2] {
				hasTwoPairs = true
			}
		}
	}
	for i:= 0; i < len(s) - 2; i++ {
		if s[i] == s[i+2] {
			return hasTwoPairs
		}
	}
	return false
}

