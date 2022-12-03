package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solve(input string, getLineChecksum func([]int) int) string {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		row := strings.Split(line, "\t")

		digits := make([]int, len(row))
		for digitIndex, digitString := range row {
			digit, err := strconv.Atoi(digitString)
			if err != nil {
				panic(err)
			}
			digits[digitIndex] = digit
		}

		sum += getLineChecksum(digits)
	}

	return fmt.Sprintf("%d", sum)
}

func getDifferenceChecksum(digits []int) int {
	sort.Ints(digits)
	return digits[len(digits)-1] - digits[0]
}

func getDivisionChecksum(digits []int) int {
	for digitIndex, digit := range digits {
		for otherDigitIndex, otherDigit := range digits {
			if digitIndex == otherDigitIndex {
				continue
			}
			if digit > otherDigit && digit%otherDigit == 0 {
				return digit / otherDigit
			}
		}
	}
	panic("No possible division")
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Input: %s\n", solve(input, getDifferenceChecksum))

	fmt.Println()

	fmt.Println("Part 2:")
	fmt.Printf("Input: %s\n", solve(input, getDivisionChecksum))
}
