package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulate(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	instructions := parseInstructions(input)
	count := simulate(instructions)
	assert.Equal(t, 13, count)
}

func TestSimulate2(t *testing.T) {
	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

	instructions := parseInstructions(input)
	count := simulate2(instructions)
	assert.Equal(t, 36, count)
}
