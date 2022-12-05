package main

import (
	"fmt"
	"strings"

	"github.com/srowles/aoc2022"
)

func main() {
	do()
}

type pair struct {
	elf1 []int64
	elf2 []int64
}

func (p pair) contained() bool {
	longest, shortest := longShort(p)
	if shortest[0] >= longest[0] && shortest[1] <= longest[1] {
		return true
	}

	return false
}

func longShort(p pair) ([]int64, []int64) {
	longest := p.elf1
	shortest := p.elf2
	if p.elf2[1]-p.elf2[0] > p.elf1[1]-p.elf1[0] {
		longest = p.elf2
		shortest = p.elf1
	}
	return longest, shortest
}

func (p pair) overlap() bool {
	longest, shortest := longShort(p)
	if shortest[0] >= longest[0] && shortest[0] <= longest[1] {
		return true
	}
	if shortest[1] >= longest[0] && shortest[1] <= longest[1] {
		return true
	}
	return false
}

func expand(elf string) []int64 {
	var s, e int64
	fmt.Sscanf(elf, "%d-%d", &s, &e)
	return []int64{s, e}
}

func do() {
	input := aoc2022.InputFromWebsite("4")
	pairs := aoc2022.Slice(input, "\n", func(l string) pair {
		e1, e2, _ := strings.Cut(l, ",")
		return pair{
			elf1: expand(e1),
			elf2: expand(e2),
		}
	})

	contained := 0
	overlap := 0
	for _, p := range pairs {
		if p.contained() {
			contained++
		}
		if p.overlap() {
			overlap++
		}
	}

	fmt.Println(contained)
	fmt.Println(overlap)
}
