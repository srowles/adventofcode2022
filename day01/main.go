package main

import "github.com/srowles/aoc2022"

func main() {
	do1()
}

func do1() {
	input := aoc2022.InputFromWebsite("1")
	rows := aoc2022.Slice[string](input, "\n", func(s string) string {
		return ""
	})
}
