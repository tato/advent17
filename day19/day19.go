package main

import (
	_ "embed"
	"fmt"
	"strings"
	"bytes"
)

//go:embed input.txt
var input string
//go:embed exa01.txt
var exa01 string

func main() {
	input := solve(parse(input))
	exa01 := solve(parse(exa01))

	fmt.Printf("Part 1 Example: %s\n", exa01.letters)
	fmt.Printf("Part 1: %s\n", input.letters)
	
	fmt.Printf("Part 2 Example: %d\n", exa01.steps)
	fmt.Printf("Part 2: %d\n", input.steps)
}

type grid struct {
	cells []byte
	w int
	h int
}

func parse(input string) grid {
	w := 0
	h := 0
	cells := make([]byte, 0)
	for _, line := range strings.Split(strings.Trim(input, "\r\n\t"), "\n") {
		row := strings.Trim(line, "\r\n\t")
		cells = append(cells, row...)

		if w != 0 && w != len(row) {
			panic("grid width is inconsistent")
		}

		w = len(row)
		h++
	}

	return grid{ cells, w, h }
}

func get(g grid, x int, y int) byte {
	if x < 0 || x >= g.w || y < 0 || y >= g.h {
		return ' '
	}
	return g.cells[y * g.w + x]
}

type result struct {
	letters string
	steps int
}

func solve(input grid) result {
	x := bytes.IndexByte(input.cells, '|')
	y := 0
	dx := 0
	dy := 1

	letters := make([]byte, 0)
	steps := 0

	for true {
		cell := get(input, x, y)
		
		if cell == ' ' {
			break
		}
	
		if cell == '+' {
			if dx == 0 && get(input, x + 1, y) != ' ' {
				dx = 1
				dy = 0
			} else if dx == 0 && get(input, x - 1, y) != ' ' {
				dx = -1
				dy = 0
			} else if dy == 0 && get(input, x, y + 1) != ' ' {
				dx = 0
				dy = 1
			} else if dy == 0 && get(input, x, y - 1) != ' ' {
				dx = 0
				dy = -1
			} else {
				panic(fmt.Sprintf("wrong turn at (%d, %d), (%d, %d)", x, y, dx, dy))
			}
		} else if cell != '|' && cell != '-' {
			letters = append(letters, cell)
		}

		x += dx
		y += dy
		steps++
	}

	return result { string(letters), steps }
}