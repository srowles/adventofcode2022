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
	dir := parse(aoc2022.InputFromWebsite("7"))
	fmt.Println(dir.sizesLessThan100000(0))
}

func do2() {
	dir := parse(aoc2022.InputFromWebsite("7"))
	freeSpace := 70000000 - dir.size()
	extraFreeSpaceRequired := 30000000 - freeSpace

	smallest = math.MaxInt64
	dir.smallestDelete(extraFreeSpaceRequired)
	fmt.Println(smallest)
}

func (d *dir) smallestDelete(size int) {
	for _, sub := range d.subdirs {
		s := sub.size()
		if s > size {
			smallest = aoc2022.Min(smallest, s)
		}
		sub.smallestDelete(size)
	}
}

var smallest int

type dir struct {
	name    string
	parent  *dir
	subdirs []*dir
	files   []file
}

func (d *dir) sizesLessThan100000(size int) int {
	mySize := d.size()
	if mySize < 100000 {
		size += mySize
	}
	for _, d := range d.subdirs {
		size += d.sizesLessThan100000(0)
	}

	return size
}

func (d *dir) print() {
	fmt.Println(d.size(), d.name)
	for _, file := range d.files {
		fmt.Println("-", file.name, file.size)
	}
	for _, d := range d.subdirs {
		d.print()
	}
}

func (d *dir) size() int {
	size := 0
	for _, file := range d.files {
		size += int(file.size)
	}
	for _, d := range d.subdirs {
		size += d.size()
	}

	return size
}

type file struct {
	name string
	size int64
}

func parse(input string) *dir {
	current := &dir{name: "/"}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "$") {
			cmd, name, _ := strings.Cut(line[2:], " ")
			switch cmd {
			case "ls":
				continue
			case "cd":
				if name == ".." {
					current = current.parent
					continue
				}
				if name == "/" {
					continue
				}
				current = byName(current, name)
				continue
			default:
				panic("unknown command " + cmd)
			}
		}
		size, name, _ := strings.Cut(line, " ")
		if size == "dir" {
			d := &dir{name: name, parent: current}
			current.subdirs = append(current.subdirs, d)
			continue
		}
		current.files = append(current.files, file{name: name, size: aoc2022.Int(size)})
	}

	for current.parent != nil {
		current = current.parent
	}
	return current
}

func byName(current *dir, name string) *dir {
	for _, d := range current.subdirs {
		if d.name == name {
			return d
		}
	}
	panic(fmt.Sprintf("not found %#v - %s", *current, name))
}
