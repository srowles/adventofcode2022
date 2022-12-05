package main

import (
	"fmt"
	"strings"

	"github.com/srowles/aoc2022"
)

func main() {
	fmt.Println("part1")
	do1()
	fmt.Println("part2")
	do2()
}

type command struct {
	count, from, to int64
}

func parseInstructions(input string) ([]aoc2022.RuneStack, []command) {
	var instructions []command
	stacks := make([]aoc2022.RuneStack, 10)
	for i := 0; i < 10; i++ {
		stacks[i] = aoc2022.NewRuneStack()
	}
	stackString, instructionString, _ := strings.Cut(input, "\n\n")
	for _, line := range strings.Split(instructionString, "\n") {
		var cmd command
		fmt.Sscanf(line, "move %d from %d to %d", &cmd.count, &cmd.from, &cmd.to)
		instructions = append(instructions, cmd)
	}

	stackLines := strings.Split(stackString, "\n")
	for i := 1; i < len(stackLines[0]); i += 4 {
		for k := len(stackLines) - 1; k >= 0; k-- {
			if stackLines[k][i] >= 65 && stackLines[k][i] <= 90 {
				stacks[(i/4)+1].Push(rune(stackLines[k][i]))
			}
		}
	}

	return stacks, instructions
}

func do1() {
	crates, instructions := parseInstructions(aoc2022.InputFromWebsite("5"))

	for _, i := range instructions {
		for k := 0; k < int(i.count); k++ {
			c := crates[i.from].Pop()
			if c != '_' {
				crates[i.to].Push(c)
			}
		}
	}

	for i := 1; i < 10; i++ {
		fmt.Println(string(crates[i].Pop()))
	}
}

func do2() {
	crates, instructions := parseInstructions(aoc2022.InputFromWebsite("5"))

	for _, i := range instructions {
		tmp := aoc2022.NewRuneStack()
		for k := 0; k < int(i.count); k++ {
			c := crates[i.from].Pop()
			if c != '_' {
				tmp.Push(c)
			}
		}
		for !tmp.IsEmpty() {
			crates[i.to].Push(tmp.Pop())
		}
	}

	for i := 1; i < 10; i++ {
		fmt.Println(string(crates[i].Pop()))
	}
}
