package main

import (
	_ "embed"
	"fmt"
)

func solve(input string, getFn func(string, int) byte) string {
	sum := 0
	for charIndex, char := range input {
		digit := int(char - '0')
		nextDigit := int(getFn(input, charIndex) - '0')
		if digit == nextDigit {
			sum += digit
		}
	}
	return fmt.Sprintf("%d", sum)
}

//go:embed input.txt
var input string

func getNext(input string, i int) byte {
	return input[(i+1)%len(input)]
}

func getAround(input string, i int) byte {
	return input[(i+len(input)/2)%len(input)]
}

func runExamples(getFn func(string, int) byte) {
	examples := []string{
		"1122",
		"1111",
		"1234",
		"91212129",
	}
	for exampleIndex, example := range examples {
		fmt.Printf("Example %d: %s\n", exampleIndex, solve(example, getFn))
	}
	fmt.Printf("Input: %s\n", solve(input, getFn))
}

func main() {
	fmt.Println("Part 1:")
	runExamples(getNext)

	fmt.Println()

	fmt.Println("Part 2:")
	runExamples(getAround)
}
