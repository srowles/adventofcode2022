package main

import (
	"fmt"

	"github.com/srowles/aoc2022"
)

func main() {
	do1()
	do2()
}

func do1() {
	forest := parse(aoc2022.InputFromWebsite("8"))
	fmt.Println(visibleTrees(forest))
}

func do2() {
	forest := parse(aoc2022.InputFromWebsite("8"))
	fmt.Println(scenicScore(forest))
}

func parse(input string) map[aoc2022.Coord]int {
	rows := aoc2022.Slice(input, "\n", func(s string) []int {
		var row []int
		for _, r := range s {
			row = append(row, int(r)-int('0'))
		}
		return row
	})

	forest := make(map[aoc2022.Coord]int)
	var y int
	for _, row := range rows {
		x := 0
		for _, h := range row {
			forest[aoc2022.Coord{X: x, Y: y}] = h
			x++
		}
		y++
	}

	return forest
}

func print(forest map[aoc2022.Coord]int) {
	var maxx, maxy int
	for c := range forest {
		maxx = aoc2022.Max(maxx, c.X)
		maxy = aoc2022.Max(maxy, c.Y)
	}
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			c := aoc2022.Coord{X: x, Y: y}
			fmt.Print(forest[c], ",")
		}
		fmt.Println()
	}
}

var dirs = []aoc2022.Coord{
	{X: -1, Y: 0},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

func visibleTrees(forest map[aoc2022.Coord]int) int {
	visible := make(map[aoc2022.Coord]bool, len(forest))
	for c := range forest {
		for _, d := range dirs {
			if !hidden(forest, c, d) {
				visible[c] = true
				break
			}
		}
	}

	return len(visible)
}

func hidden(forest map[aoc2022.Coord]int, start aoc2022.Coord, direction aoc2022.Coord) bool {
	loc := start.Move(direction)
	height := forest[start]
	for {
		otherHight, ok := forest[loc]
		if !ok {
			break
		}
		if otherHight >= height {
			return true
		}
		loc = loc.Move(direction)
	}

	return false
}

func scenicScore(forest map[aoc2022.Coord]int) int {
	top := -1
	for c := range forest {
		score := 1
		for _, d := range dirs {
			s := scenic(forest, c, d)
			score *= s
		}
		top = aoc2022.Max(top, score)
	}

	return top
}

func scenic(forest map[aoc2022.Coord]int, start aoc2022.Coord, direction aoc2022.Coord) int {
	loc := start.Move(direction)
	height := forest[start]
	for i := 0; ; i++ {
		otherHight, ok := forest[loc]
		if !ok {
			return i
		}
		if otherHight >= height {
			return i + 1
		}

		loc = loc.Move(direction)
	}
}
