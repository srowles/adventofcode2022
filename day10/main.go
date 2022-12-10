package main

import (
	"fmt"
	"strings"

	"github.com/srowles/aoc2022"
)

func main() {
	do1()
	do2()
}

func do1() {
	computer := parseInstructions(aoc2022.InputFromWebsite("10"))
	run(computer)
}

func run(computer *machine) {
	v20 := computer.run(20) * 20
	fmt.Println(v20)
	v60 := computer.run(40) * 60
	fmt.Println(v60)
	v100 := computer.run(40) * 100
	fmt.Println(v100)
	v140 := computer.run(40) * 140
	fmt.Println(v140)
	v180 := computer.run(40) * 180
	fmt.Println(v180)
	v220 := computer.run(40) * 220
	fmt.Println(v220)

	fmt.Println(v20 + v60 + v100 + v140 + v180 + v220)
}

func draw(computer *machine) {
	scanx := 0
	for c := 0; c < 240; c++ {
		spritePos := computer.x
		if scanx >= int(spritePos)-1 && scanx <= int(spritePos)+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		scanx++
		if scanx%40 == 0 {
			fmt.Println()
			scanx = 0
		}
		computer.run(1)

	}
}

func do2() {
	computer := parseInstructions(aoc2022.InputFromWebsite("10"))
	draw(computer)
}

func parseInstructions(input string) *machine {
	return &machine{
		instructions: aoc2022.Slice(input, "\n", func(s string) instruction {
			cmd, v, _ := strings.Cut(s, " ")
			if cmd == noop {
				return instruction{
					cmd: cmd,
				}
			}
			return instruction{
				cmd:   cmd,
				value: int(aoc2022.Int(v)),
			}
		}),
		x:        1,
		pipeline: make(chan instruction, 10000),
	}
}

const (
	noop = "noop"
	addx = "addx"
)

type instruction struct {
	cmd   string
	value int
}

type machine struct {
	instructions []instruction
	ip           int
	x            int64
	pipeline     chan instruction
}

func (m *machine) run(cycles int) int64 {
	var result int64
	for c := 0; c < cycles; c++ {
		if m.ip < len(m.instructions) {
			i := m.instructions[m.ip]
			if i.cmd == addx {
				// add buffer because addx takes 2 instructions to run
				m.pipeline <- instruction{cmd: noop}
			}
			m.pipeline <- i
		}
		m.ip++

		result = m.x
		cmd := <-m.pipeline
		switch cmd.cmd {
		case addx:
			m.x += int64(cmd.value)
		}
	}

	return result
}
