package knothash

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type KnotHasher struct {
	lengths  []uint8
	Numbers  []uint8
	position int
	skipSize int
}

func (h *KnotHasher) Round() {

	for _, length := range h.lengths {

		for i := 0; i < int(length)/2; i++ {
			aIndex := (h.position + i) % len(h.Numbers)
			bIndex := (h.position + int(length) - 1 - i) % len(h.Numbers)
			h.Numbers[aIndex], h.Numbers[bIndex] = h.Numbers[bIndex], h.Numbers[aIndex]
		}

		h.position = (h.position + int(length) + h.skipSize) % len(h.Numbers)

		h.skipSize++

	}
}

func (h *KnotHasher) HashIntoArray() [16]uint8 {

	h.lengths = append(h.lengths, 17, 31, 73, 47, 23)

	for roundIndex := 0; roundIndex < 64; roundIndex++ {
		h.Round()
	}

	return makeDense(h.Numbers)
}

func (h *KnotHasher) Hash() string {
	return formatDense(h.HashIntoArray())
}

func makeDense(numbers []uint8) [16]uint8 {

	var dense [16]uint8

	for i := range dense {
		elem := uint8(0)
		for j := 0; j < 16; j++ {
			elem = elem ^ numbers[i*16+j]
		}
		dense[i] = elem
	}

	return dense
}

func formatDense(dense [16]uint8) string {

	var hex bytes.Buffer

	for _, d := range dense {
		hex.WriteString(fmt.Sprintf("%x", d))
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

func FromNumbers(input string) KnotHasher {

	split := strings.Split(input, ",")
	lengths := make([]uint8, len(split))

	for index, s := range split {
		i, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			panic(true)
		}
		lengths[index] = uint8(i)
	}

	return KnotHasher{lengths: lengths, Numbers: initialNumbers()}
}

func FromAscii(input string) KnotHasher {
	lengths := make([]uint8, len(input))
	for i := range lengths {
		lengths[i] = uint8(input[i])
	}
	return KnotHasher{lengths: lengths, Numbers: initialNumbers()}
}
