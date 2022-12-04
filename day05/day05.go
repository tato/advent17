package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func solve(input string, modifyFunc func(*int)) int {
	steps := 0

	offsets := parse(input)
	ip := 0
	for ip >= 0 && ip < len(offsets) {
		inst := offsets[ip]
		modifyFunc(&offsets[ip])
		ip += inst
		steps += 1
	}

	return steps
}

func modifyBasic(inst *int) {
	*inst += 1
}

func modifyExtra(inst *int) {
	if *inst >= 3 {
		*inst -= 1
	} else {
		*inst += 1
	}
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
	fmt.Printf("%d, %d\n", solve(input, modifyBasic), solve(input, modifyExtra))
}
