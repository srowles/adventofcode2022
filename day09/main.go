package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/srowles/aoc2022"
)

func main() {
	do1()
	do2()
}

func do1() {
	grid := parseInstructions(aoc2022.InputFromWebsite("9"))
	fmt.Println(simulate(grid))
}

func do2() {
	grid := parseInstructions(aoc2022.InputFromWebsite("9"))
	fmt.Println(simulate2(grid))
}

type instruction struct {
	dir  aoc2022.Coord
	dist int
}

var dirs = map[string]aoc2022.Coord{
	"U": {X: 0, Y: -1},
	"D": {X: 0, Y: 1},
	"L": {X: -1, Y: 0},
	"R": {X: 1, Y: 0},
}

func parseInstructions(input string) []instruction {
	var instrs []instruction
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		d, l, _ := strings.Cut(line, " ")
		instrs = append(instrs, instruction{
			dir:  dirs[d],
			dist: int(aoc2022.Int(l)),
		})
	}
	return instrs
}

func simulate(instructions []instruction) int {
	tailTracker := make(map[aoc2022.Coord]bool)
	var head, tail aoc2022.Coord
	tailTracker[tail] = true
	for _, inst := range instructions {
		for i := 0; i < inst.dist; i++ {
			head = head.Move(inst.dir)
			if distance(head, tail) > 1 {
				tail = move(tail, head)
			}
			tailTracker[tail] = true
		}
	}
	return len(tailTracker)
}

func simulate2(instructions []instruction) int {
	tailTracker := make(map[aoc2022.Coord]bool)
	segments := make([]aoc2022.Coord, 10)
	tailTracker[segments[0]] = true
	for _, inst := range instructions {
		for i := 0; i < inst.dist; i++ {
			segments[0] = segments[0].Move(inst.dir)
			for idx := 0; idx < 9; idx++ {
				if distance(segments[idx], segments[idx+1]) > 1 {
					segments[idx+1] = move(segments[idx+1], segments[idx])
				}
			}
			tailTracker[segments[9]] = true
		}

	}
	return len(tailTracker)
}

func distance(tail, head aoc2022.Coord) int {
	diff := head.Diff(tail)
	return int(math.Sqrt(float64((diff.X * diff.X) + (diff.Y * diff.Y))))
}

func move(tail, head aoc2022.Coord) aoc2022.Coord {
	diff := head.Diff(tail)
	if aoc2022.Abs(diff.X) > 0 {
		diff.X = diff.X / aoc2022.Abs(diff.X)
	}
	if aoc2022.Abs(diff.Y) > 0 {
		diff.Y = diff.Y / aoc2022.Abs(diff.Y)
	}
	return tail.Move(diff)
}
