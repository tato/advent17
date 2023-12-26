package main

import (
	_ "embed"
	"fmt"
	"strings"
	"strconv"
)

//go:embed input.txt
var input string
//go:embed exa01.txt
var exa01 string
//go:embed exa02.txt
var exa02 string

func main() {
	input := parse(input)
	exa01 := parse(exa01)
	exa02 := parse(exa02)

	fmt.Printf("Part 1 Example: %d\n", part1(exa01))
	fmt.Printf("Part 1: %d\n", part1(input))

	fmt.Printf("Part 2 Example: %d\n", part2(exa02))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type particle struct {
	p [3]int64
	v [3]int64
	a [3]int64
}

func parse(input string) []particle {
	particles := make([]particle, 0)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(strings.TrimSpace(line), ", ")
		p := parseVector(parts[0])
		v := parseVector(parts[1])
		a := parseVector(parts[2])
		particles = append(particles, particle{ p, v, a })
	}
	return particles
}

func parseVector(input string) [3]int64 {
	parts := strings.Split(input[3:len(input)-1], ",")
	x, _ := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 0)
	y, _ := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 0)
	z, _ := strconv.ParseInt(strings.TrimSpace(parts[2]), 10, 0)
	return [3]int64{ x, y, z }
}

func abs(n int64) int64 {
	if (n < 0) {
		return -n
	}
	return n
}
func length(v [3]int64) int64 {
	return abs(v[0]) + abs(v[1]) + abs(v[2])
}
func add(a [3]int64, b [3]int64) [3]int64 {
	return [3]int64{ a[0] + b[0], a[1] + b[1], a[2] + b[2] }
}
func sub(a [3]int64, b [3]int64) [3]int64 {
	return [3]int64{ a[0] - b[0], a[1] - b[1], a[2] - b[2] }
}
func equals(a [3]int64, b [3]int64) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}

func part1(input []particle) int {
	
	minIndex := 0
	minVal := length(input[0].a)
	for index, particle := range input {
		if length(particle.a) < minVal {
			minIndex = index
			minVal = length(particle.a)
		}
	}

	return minIndex
}

func update(particle *particle) {
	particle.v = add(particle.v, particle.a)
	particle.p = add(particle.p, particle.v)
}

func part2(input []particle) int {
	particles := append(make([]particle, 0, len(input)), input...)
	
	for round := 0; round < 1000; round++ {
		for index := range particles {
			update(&particles[index])
		}

		positions := make(map[[3]int64]int)
		for _, particle := range particles {
			positions[particle.p]++
		}

		aliveParticles := make([]particle, 0, len(particles))
		for _, particle := range particles {
			if positions[particle.p] == 1 {
				aliveParticles = append(aliveParticles, particle)
			}
		}
		particles = aliveParticles
	}
	
	return len(particles)
}