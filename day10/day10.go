package main

import (
	_ "embed"
	"fmt"

	"supok.es/advent17/day10/knothash"
)

//go:embed input.txt
var input string

func main() {

	numbersHasher := knothash.FromNumbers(input)
	numbersHasher.Round()
	part1 := int(numbersHasher.Numbers[0]) * int(numbersHasher.Numbers[1])

	asciiHasher := knothash.FromAscii(input)
	hashResult := asciiHasher.Hash()

	fmt.Printf("Part 1: %d, Part 2: %s\n", part1, hashResult)
}
