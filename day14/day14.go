package main

import (
	"fmt"

	"supok.es/advent17/day10/knothash"
)

func solve(input string) {

	usedSquares := 0

	var disk [128][128]bool

	for index := 0; index < 128; index++ {
		asciiHasher := knothash.FromAscii(fmt.Sprintf("%s-%d", input, index))
		hashResult := asciiHasher.HashIntoArray()
		for numberIndex, number := range hashResult {
			for bitIndex := 0; bitIndex < 8; bitIndex++ {
				if (number>>(8-bitIndex-1))&1 == 1 {
					disk[index][numberIndex*8+bitIndex] = true
					usedSquares++
				}
			}
		}
	}

	var regionsCount uint

	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			if disk[x][y] {
				regionsCount++
				adjacents := [][2]int{{x, y}}
				for len(adjacents) > 0 {
					position := adjacents[len(adjacents)-1]
					adjacents = adjacents[0 : len(adjacents)-1]

					disk[position[0]][position[1]] = false

					for _, direction := range directions {
						dx := position[0] + direction[0]
						dy := position[1] + direction[1]
						if dx >= 0 && dy >= 0 && dx < len(disk[0]) && dy < len(disk) && disk[dx][dy] {
							adjacents = append(adjacents, [2]int{dx, dy})
						}
					}
				}
			}
		}
	}

	fmt.Printf("Used Squares: %d, Regions Count: %d\n", usedSquares, regionsCount)
}

var directions [][2]int = [][2]int{{+1, 0}, {-1, 0}, {0, +1}, {0, -1}}

func main() {
	solve("vbqugkhl")
}
