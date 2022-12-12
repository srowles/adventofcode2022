package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	monkeys, _ := parseMonkeys(ex)
	for i := 0; i < 20; i++ {
		processItems(monkeys, divideByThree)
	}
	printMonkeys(monkeys)
	t.FailNow()
}

func TestParse2(t *testing.T) {
	monkeys, factor := parseMonkeys(ex)
	for i := 0; i < 10000; i++ {
		processItems(monkeys, func(u uint64) uint64 { return u % factor })
	}
	printMonkeys(monkeys)
	t.FailNow()
}

var ex = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`
