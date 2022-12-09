package main

import "fmt"

func solve(input int) {
	calm := spinlockCalm(input, 2017)
	angry := spinlockAngry(input, 50_000_000)

	fmt.Printf("After: %d, Angry: %d\n", calm, angry)
}

func spinlockCalm(steps int, iterations int) int {
	buffer := []int{0}
	position := 0

	for nextNumber := 1; nextNumber <= iterations; nextNumber++ {
		position = (position + steps + 1) % len(buffer)
		newBuffer := make([]int, 0, len(buffer)+1)
		for index, value := range buffer {
			newBuffer = append(newBuffer, value)
			if index == position {
				newBuffer = append(newBuffer, nextNumber)
			}
		}
		buffer = newBuffer
	}

	return buffer[(position+2)%len(buffer)]
}

func spinlockAngry(steps int, iterations int) int {
	position := 0
	result := 0
	for bufferLength := 1; bufferLength <= iterations; bufferLength++ {
		position = (position + steps + 1) % bufferLength
		if position == 0 {
			result = bufferLength
		}
	}
	return result
}

func main() {
	solve(316)
}
