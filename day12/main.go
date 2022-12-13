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
	grid, start, end := parseGrid(aoc2022.InputFromWebsite("12"))
	fmt.Println(walk(grid, start, end))
}

func do2() {
	grid, _, end := parseGrid(aoc2022.InputFromWebsite("12"))
	var as []aoc2022.Coord
	for c, r := range grid {
		if r == 'a' {
			as = append(as, c)
		}
	}

	shortest := math.MaxInt
	for _, a := range as {
		dist := walk(grid, a, end)
		if dist != -1 {
			shortest = aoc2022.Min(shortest, dist)
		}
	}
	fmt.Println(shortest)
}

func walk(grid map[aoc2022.Coord]rune, start, end aoc2022.Coord) int {
	distances := make(map[aoc2022.Coord]int)
	next := []aoc2022.Coord{start}
	for len(next) > 0 {
		pos := next[0]
		for _, n := range accessible(pos, grid) {
			if _, ok := distances[n]; ok {
				// already visited, skip
				continue
			}
			distances[n] = distances[pos] + 1
			next = append(next, n)
		}

		next = next[1:]
	}

	dist, ok := distances[end]
	if !ok {
		return -1
	}
	return dist
}

var moves = []aoc2022.Coord{
	{X: 1, Y: 0},
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 0, Y: -1},
}

func accessible(p aoc2022.Coord, grid map[aoc2022.Coord]rune) []aoc2022.Coord {
	var res []aoc2022.Coord
	for _, m := range moves {
		to := p.Move(m)
		if _, ok := grid[to]; !ok {
			continue
		}
		cost := grid[to] - grid[p]
		if cost <= 1 {
			res = append(res, to)
		}
	}
	return res
}

func parseGrid(input string) (map[aoc2022.Coord]rune, aoc2022.Coord, aoc2022.Coord) {
	var start, end aoc2022.Coord
	grid := make(map[aoc2022.Coord]rune)
	for row, line := range strings.Split(input, "\n") {
		for x, r := range line {
			c := aoc2022.Coord{X: x, Y: row}
			if r == 'E' {
				r = 'z'
				end = c
			}
			if r == 'S' {
				r = 'a'
				start = c
			}
			grid[c] = r
		}
	}

	return grid, start, end
}

func printGrid(grid map[aoc2022.Coord]rune) {
	var maxx, maxy int
	for c := range grid {
		maxx = aoc2022.Max(maxx, c.X)
		maxy = aoc2022.Max(maxy, c.Y)
	}

	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			c := aoc2022.Coord{X: x, Y: y}
			fmt.Print(string(grid[c]))
		}
		fmt.Println()
	}
}
