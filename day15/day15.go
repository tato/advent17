package main

import "fmt"

var factorA uint64 = 16807
var factorB uint64 = 48271
var divisor uint64 = 2147483647
var mask uint64 = 0xFFFF

func solve(input pair) {
	basic := judgeBasic(input)
	extra := judgeExtra(input)
	fmt.Printf("Basic: %d, Extra %d\n", basic, extra)
}

type pair struct{ a, b uint64 }

func judgeBasic(input pair) uint {
	var counter uint
	for index := 0; index <= 40_000_000; index++ {
		input.a = (input.a * factorA) % divisor
		input.b = (input.b * factorB) % divisor
		if input.a&mask == input.b&mask {
			counter++
		}
	}
	return counter
}

func judgeExtra(input pair) uint {
	var counter uint
	for index := 0; index <= 5_000_000; index++ {
		input.a = (input.a * factorA) % divisor
		for input.a%4 != 0 {
			input.a = (input.a * factorA) % divisor
		}
		input.b = (input.b * factorB) % divisor
		for input.b%8 != 0 {
			input.b = (input.b * factorB) % divisor
		}
		if input.a&mask == input.b&mask {
			counter++
		}
	}
	return counter
}

var myInput = pair{289, 629}

func main() {
	solve(myInput)
}
