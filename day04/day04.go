package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func solve(input string) (int, int) {
	var countWords int
	var countAnagrams int

	for _, line := range strings.Split(input, "\n") {
		oneIfValidWordPhrase := 1
		oneIfValidAnagramPhrase := 1

		words := make(map[string]struct{})
		anagrams := make(map[uint32]struct{})

		for _, word := range strings.Split(line, " ") {
			bitset := asBitset(word)

			if _, ok := words[word]; ok {
				oneIfValidWordPhrase = 0
			}
			if _, ok := anagrams[bitset]; ok {
				oneIfValidAnagramPhrase = 0
			}

			words[word] = struct{}{}
			anagrams[bitset] = struct{}{}
		}

		countWords += oneIfValidWordPhrase
		countAnagrams += oneIfValidAnagramPhrase
	}

	return countWords, countAnagrams
}

func asBitset(word string) uint32 {
	var bitset uint32
	for _, char := range word {
		bitset |= 1 << uint32(char-'a')
	}
	return bitset
}

//go:embed input.txt
var input string

func main() {
	one, two := solve(input)
	fmt.Printf("One %d, Two %d\n", one, two)
}
