package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/srowles/aoc2022"
)

func main() {
	do1()
	do2()
}

var divideByThree = func(u uint64) uint64 { return u / 3 }

func do1() {
	monkeys, _ := parseMonkeys(aoc2022.InputFromWebsite("11"))

	for i := 0; i < 20; i++ {
		processItems(monkeys, divideByThree)
	}
	printMonkeys(monkeys)
}

func do2() {
	monkeys, factor := parseMonkeys(aoc2022.InputFromWebsite("11"))
	for i := 0; i < 10000; i++ {
		processItems(monkeys, func(u uint64) uint64 { return u % factor })
	}
	printMonkeys(monkeys)
}

func processItems(monkeys []*monkey, relief func(uint64) uint64) {
	for _, m := range monkeys {
		for _, i := range m.items {
			m.inspectionCount++
			i = m.op(i)
			i = relief(i)
			var target int
			if m.test(i) {
				target = m.t
			} else {
				target = m.f
			}
			monkeys[target].items = append(monkeys[target].items, i)
		}
		m.items = []uint64{}
	}
}

type monkey struct {
	t               int
	f               int
	op              func(uint64) uint64
	test            func(uint64) bool
	items           []uint64
	inspectionCount uint64
}

func printMonkeys(monkeys []*monkey) {
	var scores []int
	for _, m := range monkeys {
		scores = append(scores, int(m.inspectionCount))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	fmt.Println(scores)
	fmt.Println(scores[0] * scores[1])
}

func parseMonkeys(input string) ([]*monkey, uint64) {
	factor := uint64(1)
	monkeys := aoc2022.Slice(input, "\n\n", func(s string) *monkey {
		lines := strings.Split(s, "\n")
		_, itemsStr, _ := strings.Cut(lines[1], ": ")
		items := aoc2022.Slice(itemsStr, ", ", func(s string) uint64 { return aoc2022.UInt(s) })
		_, opStr, _ := strings.Cut(lines[2], "= old ")
		_, testStr, _ := strings.Cut(lines[3], "by ")
		factor *= aoc2022.UInt(testStr)
		_, tStr, _ := strings.Cut(lines[4], "monkey ")
		_, fStr, _ := strings.Cut(lines[5], "monkey ")
		return &monkey{
			items: items,
			op:    opFunc(opStr),
			test:  func(i uint64) bool { return i%aoc2022.UInt(testStr) == 0 },
			t:     int(aoc2022.UInt(tStr)),
			f:     int(aoc2022.UInt(fStr)),
		}
	})

	return monkeys, factor
}

func opFunc(input string) func(uint64) uint64 {
	if input == "* old" {
		return func(i uint64) uint64 {
			return i * i
		}
	}

	op, val, _ := strings.Cut(input, " ")
	if op == "*" {
		return func(i uint64) uint64 {
			return i * aoc2022.UInt(val)
		}
	}

	return func(i uint64) uint64 {
		return i + aoc2022.UInt(val)
	}
}
