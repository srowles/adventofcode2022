package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	grid, start, end := parseGrid(ex)
	printGrid(grid)
	fmt.Println(walk(grid, start, end))
	t.FailNow()
}

var ex = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
