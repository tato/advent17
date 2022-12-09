package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func solve(input string) {
	programs := []byte("abcdefghijklmnop")

	programs = dance(programs, input)
	programsCopy := string(programs)

	one_billion := 1_000_000_000
	period := 0

	for i := 0; i < one_billion; i++ {
		programs = dance(programs, input)
		if string(programs) == "bkgcdefiholnpmja" {
			period = i
			break
		}
	}

	programs = []byte("abcdefghijklmnop")
	nth := one_billion % (period + 1)
	for i := 0; i < nth; i++ {
		programs = dance(programs, input)
	}

	fmt.Printf("Once: %s, Billion: %s\n", programsCopy, string(programs))
}

func dance(programs []byte, dance string) []byte {
	for _, move := range strings.Split(dance, ",") {
		switch move[0] {
		case 's':
			n, err := strconv.Atoi(move[1:])
			_ = err
			programs = append(programs[int(16-n):], programs[:int(16-n)]...)
		case 'x':
			split := strings.Split(move[1:], "/")
			a, aerr := strconv.Atoi(split[0])
			b, berr := strconv.Atoi(split[1])
			_, _ = aerr, berr
			programs[a], programs[b] = programs[b], programs[a]
		case 'p':
			for index, c := range programs {
				if c == move[1] {
					programs[index] = move[3]
				} else if c == move[3] {
					programs[index] = move[1]
				}
			}
		}
	}
	return programs
}

//go:embed input.txt
var myInput string

func main() {
	solve(myInput)
}
