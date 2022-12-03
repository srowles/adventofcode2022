package main

import (
	"fmt"

	"github.com/srowles/aoc2022"
)

func main() {
	do()
}

func do() {
	input := aoc2022.InputFromWebsite("3")
	backpacks := aoc2022.Slice(input, "\n", func(l string) string {
		return l
	})

	var overlaps []rune
	for _, b := range backpacks {
		mid := len(b) / 2
		maps := parseRows([]string{b[0:mid], b[mid:]})
		r := intersect(2, maps...)
		overlaps = append(overlaps, r)
	}

	fmt.Println(score(overlaps))

	var overlaps2 []rune
	i := 0
	for i < len(backpacks) {
		maps := parseRows(backpacks[i : i+3])
		r := intersect(3, maps...)
		overlaps2 = append(overlaps2, r)
		i += 3
	}

	fmt.Println(score(overlaps2))
}

func parseRows(rows []string) []map[rune]struct{} {
	var maps []map[rune]struct{}
	for _, row := range rows {
		m := make(map[rune]struct{})
		for _, a := range row {
			m[a] = struct{}{}
		}
		maps = append(maps, m)
	}

	return maps
}

func intersect(expectedCount int, maps ...map[rune]struct{}) rune {
	for r := range maps[0] {
		count := 1
		for i := 1; i < len(maps); i++ {
			for or := range maps[i] {
				if or == r {
					count++
				}
			}
		}
		if count == expectedCount {
			return r
		}
	}

	return 0
}

func score(overlaps []rune) int {
	value := 0
	for _, c := range overlaps {
		if int(c) >= int('a') && int(c) <= int('z') {
			value += int(c) - int('a') + 1
		} else {
			value += int(c) - int('A') + 27
		}
	}

	return value
}
