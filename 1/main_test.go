package main

import "testing"

type Test struct {
	input    []int
	target   int
	expected int
}

var tablesA = []Test{
	{[]int{1721, 979, 366, 299, 675, 1456}, 2020, 514579},
}

var tablesB = []Test{
	{[]int{1721, 979, 366, 299, 675, 1456}, 2020, 241861950},
}

func TestPart1(t *testing.T) {
	for _, test := range tablesA {
		sum := Part1(test.input, test.target)
		if sum != test.expected {
			t.Errorf("Part1 was incorrect, got: %d, want: %d.", sum, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	for _, test := range tablesB {
		sum := Part2(test.input, test.target)
		if sum != test.expected {
			t.Errorf("Part2 was incorrect, got: %d, want: %d.", sum, test.expected)
		}
	}
}
