package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func solve(input string) int {
	steps := 0

	offsets := parse(input)
	ip := 0
	for ip >= 0 && ip < len(offsets) {
		inst := offsets[ip]
		offsets[ip] += 1
		ip += inst
		steps += 1
	}

	return steps
}

func solveExtra(input string) int {
	steps := 0

	offsets := parse(input)
	ip := 0
	for ip >= 0 && ip < len(offsets) {
		inst := offsets[ip]
		if inst >= 3 {
			offsets[ip] -= 1
		} else {
			offsets[ip] += 1
		}
		ip += inst
		steps += 1
	}

	return steps
}

func parse(input string) []int {
	inputSplit := strings.Split(input, "\n")
	offsets := make([]int, len(inputSplit))
	for index, offset := range inputSplit {
		offset, err := strconv.Atoi(offset)
		if err != nil {
			panic(err)
		}
		offsets[index] = offset
	}
	return offsets
}

//go:embed input.txt
var input string

func main() {
	fmt.Printf("%d, %d\n", solve(input), solveExtra(input))
}
