package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func solve(input string) (uint, uint) {
	var cycles uint

	memory := parse(input)

	configurations := make(map[Memory]uint)

	var repeatConfigCycle uint

	for configFound := false; !configFound; repeatConfigCycle, configFound = configurations[memory] {
		configurations[memory] = cycles

		maxIndex := 0
		for index, blocks := range memory {
			if blocks > memory[maxIndex] {
				maxIndex = index
			}
		}

		redistribute := memory[maxIndex]
		memory[maxIndex] = 0
		index := maxIndex

		for ; redistribute > 0; redistribute-- {
			index = (index + 1) % len(memory)
			memory[index] += 1
		}

		cycles++
	}

	return cycles, cycles - repeatConfigCycle
}

type Memory [16]uint

func parse(input string) Memory {
	var memory Memory
	inputSplit := strings.Split(input, "\t")
	for index, bank := range inputSplit {
		bank, err := strconv.ParseUint(bank, 10, 0)
		if err != nil {
			panic(err)
		}
		memory[index] = uint(bank)
	}
	return memory
}

//go:embed input.txt
var input string

func main() {
	cycles, period := solve(input)
	fmt.Printf("%d, %d\n", cycles, period)
}
