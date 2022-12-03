package main

import (
	"fmt"
	"strings"

	"github.com/srowles/aoc2022"
)

func main() {
	do()
}

type Instruction struct {
	R1 string
	R2 string
}

func (i Instruction) Score() int {
	score := int(i.R2[0]) - int('W')
	switch i.R2 {
	case "X": // rock
		if i.R1 == "C" {
			score += 6
		}
		if i.R1 == "A" {
			score += 3
		}
	case "Y": // paper
		if i.R1 == "A" {
			score += 6
		}
		if i.R1 == "B" {
			score += 3
		}
	case "Z": // scissors
		if i.R1 == "B" {
			score += 6
		}
		if i.R1 == "C" {
			score += 3
		}
	}

	return score
}

func (i Instruction) Score2() int {
	score := 0
	switch i.R2 {
	case "X": // lose
		if i.R1 == "A" {
			score += 3
		}
		if i.R1 == "B" {
			score += 1
		}
		if i.R1 == "C" {
			score += 2
		}
	case "Y": // draw
		score += 3
		if i.R1 == "A" {
			score += 1
		}
		if i.R1 == "B" {
			score += 2
		}
		if i.R1 == "C" {
			score += 3
		}
	case "Z": // win
		score += 6
		if i.R1 == "A" {
			score += 2
		}
		if i.R1 == "B" {
			score += 3
		}
		if i.R1 == "C" {
			score += 1
		}
	}

	return score
}

func do() {
	input := aoc2022.InputFromWebsite("2")
	instructions := aoc2022.Slice(input, "\n", func(l string) Instruction {
		r1, r2, _ := strings.Cut(l, " ")
		return Instruction{R1: r1, R2: r2}
	})

	score := 0
	score2 := 0
	for _, i := range instructions {
		score += i.Score()
		score2 += i.Score2()
	}
	fmt.Println(score)
	fmt.Println(score2)
}
