package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Parallel()
	// Config

	// Test Cases
	tests := map[string]struct {
		input    []string
		expected int
	}{
		"ex1": {
			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			expected: 18,
		},
		"ex2": {
			input: []string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			expected: 52,
		},
	}

	// Testing
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			maps := parseRows(test.input)

			var overlaps []rune
			overlaps = append(overlaps, intersect(3, maps...))
			result := score(overlaps)
			assert.Equal(t, test.expected, result)
		})
	}
}
