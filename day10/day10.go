package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type knotHasher struct {
	lengths  []uint8
	numbers  []uint8
	position int
	skipSize int
}

func (h *knotHasher) round() {

	for _, length := range h.lengths {

		for i := 0; i < int(length)/2; i++ {
			aIndex := (h.position + i) % len(h.numbers)
			bIndex := (h.position + int(length) - 1 - i) % len(h.numbers)
			h.numbers[aIndex], h.numbers[bIndex] = h.numbers[bIndex], h.numbers[aIndex]
		}

		h.position = (h.position + int(length) + h.skipSize) % len(h.numbers)

		h.skipSize++

	}
}

func (h *knotHasher) hash() string {

	h.lengths = append(h.lengths, 17, 31, 73, 47, 23)

	for roundIndex := 0; roundIndex < 64; roundIndex++ {
		h.round()
	}

	return makeDense(h.numbers)
}

func makeDense(numbers []uint8) string {

	hex := bytes.NewBuffer(make([]byte, 32))

	i := 0
	for i < len(numbers) {
		elem := uint8(0)
		for j := 0; j < 16; j++ {
			elem = elem ^ numbers[i+j]
		}
		i += 16
		hex.WriteString(fmt.Sprintf("%x", elem))
	}

	return hex.String()
}

func initialNumbers() []uint8 {
	numbers := make([]uint8, 256)
	for i := range numbers {
		numbers[i] = uint8(i)
	}
	return numbers
}

func fromNumbers(input string) knotHasher {

	split := strings.Split(input, ",")
	lengths := make([]uint8, len(split))

	for index, s := range split {
		i, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			panic(true)
		}
		lengths[index] = uint8(i)
	}

	return knotHasher{lengths: lengths, numbers: initialNumbers()}
}

func fromAscii(input string) knotHasher {
	lengths := make([]uint8, len(input))
	for i := range lengths {
		lengths[i] = uint8(input[i])
	}
	return knotHasher{lengths: lengths, numbers: initialNumbers()}
}

//go:embed input.txt
var input string

func main() {

	numbersHasher := fromNumbers(input)
	numbersHasher.round()
	part1 := int(numbersHasher.numbers[0]) * int(numbersHasher.numbers[1])

	asciiHasher := fromAscii(input)
	hashResult := asciiHasher.hash()

	fmt.Printf("Part 1: %d, Part 2: %s\n", part1, hashResult)
}
