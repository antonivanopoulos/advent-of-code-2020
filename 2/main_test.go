package main

import "testing"

type Test struct {
	input    []string
	expected int
}

var tables1 = []Test{
	{[]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}, 2},
}

var tables2 = []Test{
	{[]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}, 1},
}

func TestPart1(t *testing.T) {
	for _, test := range tables1 {
		sum := Part1(test.input)
		if sum != test.expected {
			t.Errorf("Part1 was incorrect, got: %d, want: %d.", sum, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	for _, test := range tables2 {
		sum := Part2(test.input)
		if sum != test.expected {
			t.Errorf("Part2 was incorrect, got: %d, want: %d.", sum, test.expected)
		}
	}
}
