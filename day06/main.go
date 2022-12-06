package main

import (
	"fmt"

	"github.com/srowles/aoc2022"
)

func main() {
	fmt.Println("part1")
	do1()
	fmt.Println("part2")
	do2()
}

func do1() {
	fmt.Println(markerPos(aoc2022.InputFromWebsite("6"), 4))
}

func do2() {
	fmt.Println(markerPos(aoc2022.InputFromWebsite("6"), 14))
}

func markerPos(input string, length int) int {
	for i := length; i < len(input); i++ {
		if allDifferent(input[i-length : i]) {
			return i
		}
	}
	return 0
}

func allDifferent(s string) bool {
	exists := make(map[rune]bool)
	for _, r := range s {
		if exists[r] {
			return false
		}
		exists[r] = true
	}

	return true
}
