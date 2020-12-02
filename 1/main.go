package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lines, err := ReadInputFile("./input.txt")
	if err != nil {
		panic(err)
	}

	entries := make([]int, 0, len(lines))
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		entries = append(entries, n)
	}

	fmt.Printf("Part1: %d\n", Part1(entries, 2020))
	fmt.Printf("Part2: %d\n", Part2(entries, 2020))
}

// ReadInputFile reads in the given input file, and converts each line into an integer
func ReadInputFile(path string) (lines []string, err error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines = strings.Split(string(dat), "\n")
	return lines, nil
}

// Part1 takes an input of an array of integers and a target integer, and returns the product of the
// two numbers that add together to equal the target.
func Part1(input []int, target int) int {
	for i := 0; i < len(input)-3; i++ {
		for j := i + 1; j < len(input)-2; j++ {
			if input[i]+input[j] == target {
				return input[i] * input[j]
			}
		}
	}

	return -1
}

// Part2 takes an input of an array of integers and a target integer, and returns the product of the
// three numbers that add together to equal the target.
func Part2(input []int, target int) int {
	// This'd be a slow boi in bigger data sets, but, the input dataset is small enough that
	// it's not a worry here and this is the easiest way to get our answer.
	for i := 0; i < len(input)-3; i++ {
		for j := i + 1; j < len(input)-2; j++ {
			for k := j + 1; k < len(input)-1; k++ {
				if input[i]+input[j]+input[k] == target {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}

	return -1
}
