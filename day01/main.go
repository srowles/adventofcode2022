package main

import (
	"fmt"
	"sort"

	"github.com/srowles/aoc2022"
)

func main() {
	do()
}

func do() {
	input := aoc2022.InputFromWebsite("1")
	rows := aoc2022.Slice(input, "\n", func(l string) string {
		return l
	})

	var elf int64 = 0
	var elves []int64
	for _, row := range rows {
		if row == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}

		elf += aoc2022.Int(row)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Println(elves[0])
	fmt.Println(elves[0] + elves[1] + elves[2])
}
