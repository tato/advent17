package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input := strings.TrimSpace(input)
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) int {
	return solve(input, func (input string, i int) byte {
		return input[(i+1)%len(input)]
	})
}

func part2(input string) int {
	return solve(input, func (input string, i int) byte {
		return input[(i+len(input)/2)%len(input)]
	})
}

func solve(input string, getFn func(string, int) byte) int {
	sum := 0
	for charIndex, char := range input {
		digit := int(char - '0')
		nextDigit := int(getFn(input, charIndex) - '0')
		if digit == nextDigit {
			sum += digit
		}
	}
	return sum
}
