package day_04

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Day4Part1(input string) int {
	for i := 0 ;; i++ {
		key := fmt.Sprintf("%s%d", input, i)	
		sum := fmt.Sprintf("%x", md5.Sum([]byte(key)))
		if strings.HasPrefix(sum ,"000000") {
			fmt.Printf("Attempt %d\n", i)
			fmt.Println(sum)
			return i
		}
	}
	return 0
}

