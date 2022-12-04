package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func solve(input string) (string, uint) {
	programs := parse(input)

	bottomName := strings.Split(input, " ")[0]

	for ; programs[bottomName].parent != ""; bottomName = programs[bottomName].parent {
	}

	calculateWeights(programs, bottomName)
	fixedWeight := calculateFixedWeight(programs, bottomName)

	return bottomName, fixedWeight
}

type program struct {
	weight            uint
	parent            string
	children          []string
	childrenWeights   []uint
	childrenWeightSum uint
	unbalancedBranch  *int
}

func parse(input string) map[string]program {
	tree := make(map[string]program)

	inputSplit := strings.Split(input, "\n")
	for _, programLine := range inputSplit {
		tokens := strings.Split(programLine, " ")
		programName := tokens[0]
		programWeight, err := strconv.ParseUint(strings.Trim(tokens[1], "()"), 10, 0)
		if err != nil {
			panic(err)
		}

		program := tree[programName]
		program.weight = uint(programWeight)
		program.unbalancedBranch = nil
		tree[programName] = program

		if len(tokens) <= 3 {
			continue
		}

		children := make([]string, len(tokens[3:]))

		for childIndex, token := range tokens[3:] {
			childName := strings.Trim(token, " ,")

			child := tree[childName]
			child.parent = programName
			tree[childName] = child

			children[childIndex] = childName
		}

		program = tree[programName]
		program.children = children
		program.childrenWeights = make([]uint, len(children))
		tree[programName] = program
	}

	return tree
}

func calculateWeights(tree map[string]program, name string) uint {
	program := tree[name]
	for index, child := range program.children {
		program.childrenWeights[index] = calculateWeights(tree, child)
		program.childrenWeightSum += program.childrenWeights[index]
	}
	calculateUnbalance(&program)
	tree[name] = program
	return program.childrenWeightSum + program.weight
}

func calculateUnbalance(program *program) {
	for candidateIndex, candidateWeight := range program.childrenWeights {
		diffAll := true
		for compIndex, compWeight := range program.childrenWeights {
			if compIndex == candidateIndex {
				continue
			}
			if candidateWeight == compWeight {
				diffAll = false
				break
			}
		}
		if diffAll {
			program.unbalancedBranch = &candidateIndex
			break
		}
	}
}

func calculateFixedWeight(tree map[string]program, name string) uint {
	program := tree[name]

	unbalancedName := program.children[*program.unbalancedBranch]
	unbalanced := tree[unbalancedName]

	balancedIndex := (*program.unbalancedBranch + 1) % len(program.children)
	balanced := tree[program.children[balancedIndex]]

	if unbalanced.unbalancedBranch == nil {
		return balanced.weight + balanced.childrenWeightSum - unbalanced.childrenWeightSum
	} else {
		return calculateFixedWeight(tree, unbalancedName)
	}
}

//go:embed input.txt
var input string

func main() {
	bottomProgram, fixedWeight := solve(input)
	fmt.Printf("Bottom: %s, Fixed Weight: %d\n", bottomProgram, fixedWeight)
}
