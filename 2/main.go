package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Policy reflects the policy scheme to validate a given password against
type Policy struct {
	X      int
	Y      int
	Letter rune
}

// Input is the combination of a parsed validation Policy, and the password to be validated
type Input struct {
	Policy   Policy
	Password string
}

func main() {
	lines, err := ReadInputFile("./input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part1: %d\n", Part1(lines))
	fmt.Printf("Part2: %d\n", Part2(lines))
}

// ReadInputFile reads in the given input file, and converts each line into an integer
func ReadInputFile(path string) (entries []string, err error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(dat), "\n")

	return lines, nil
}

// Part1 calculates the number of valid passwords for part one of day 2.
func Part1(entries []string) int {
	valid := 0

	inputs := ParseInputs(entries)
	for _, i := range inputs {
		if i.Policy.ValidatePasswordForRentalShop(i.Password) {
			valid++
		}
	}

	return valid
}

// Part2 calculates the number of valid passwords for part two of day 2.
func Part2(entries []string) int {
	valid := 0

	inputs := ParseInputs(entries)
	for _, i := range inputs {
		if i.Policy.ValidatePasswordForCorporate(i.Password) {
			valid++
		}
	}

	return valid
}

// ParseInputs takes an array of entries and parses them into valid Input structs
func ParseInputs(entries []string) []Input {
	inputs := make([]Input, 0)
	for _, string := range entries {
		splits := strings.FieldsFunc(string, Split)

		if len(splits) == 0 {
			continue
		}

		x, _ := strconv.Atoi(splits[0])
		y, _ := strconv.Atoi(splits[1])

		policy := Policy{
			X:      x,
			Y:      y,
			Letter: []rune(splits[2])[0],
		}

		input := Input{
			Policy:   policy,
			Password: splits[3],
		}

		inputs = append(inputs, input)
	}

	return inputs
}

// Split splits a policy / password entry across multiple runes present in the structure.
func Split(r rune) bool {
	return r == ':' || r == '-' || r == ' '
}

// ValidatePasswordForRentalShop will validate the given password against the
// password policy using the North Pole Toboggan Rental Shop validation scheme,
// where the specified letter in the policy can only appear between X or Y times.
func (p Policy) ValidatePasswordForRentalShop(password string) bool {
	t := 0
	for _, c := range password {
		if c == p.Letter {
			t++
		}
	}

	return t >= p.X && t <= p.Y
}

// ValidatePasswordForCorporate will validate the given password against the
// password policy using the Official Toboggan Corporate Authentication System
// validation scheme, where the specified letter can only be present at one of
// X and Y positions specified in the policy.
func (p Policy) ValidatePasswordForCorporate(password string) bool {
	x := rune(password[p.X-1]) == p.Letter
	y := rune(password[p.Y-1]) == p.Letter

	return (x || y) && !(x && y)
}
