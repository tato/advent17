package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func solve(input string) (maximum int, finalMaximum int) {
	maximum = math.MinInt
	registers := make(map[string]int)

	code := strings.Split(input, "\n")

	for _, instr := range code {
		instr := parseInstr(instr)
		if evaluateCondition(registers, instr) {
			apply(registers, instr)
			if registers[instr.register] > maximum {
				maximum = registers[instr.register]
			}
		}
	}

	finalMaximum = math.MinInt
	for _, v := range registers {
		if v > finalMaximum {
			finalMaximum = v
		}
	}

	return
}

type instruction struct {
	register          string
	operation         string
	operand           int
	condition         string
	conditionRegister string
	conditionOperand  int
}

func parseInstr(instr string) instruction {
	parts := strings.Split(instr, " ")
	return instruction{
		register:          parts[0],
		operation:         parts[1],
		operand:           parseInt(parts[2]),
		condition:         parts[5],
		conditionRegister: parts[4],
		conditionOperand:  parseInt(parts[6]),
	}
}

func evaluateCondition(registers map[string]int, instr instruction) bool {
	switch instr.condition {
	case ">":
		return registers[instr.conditionRegister] > instr.conditionOperand
	case "<":
		return registers[instr.conditionRegister] < instr.conditionOperand
	case ">=":
		return registers[instr.conditionRegister] >= instr.conditionOperand
	case "<=":
		return registers[instr.conditionRegister] <= instr.conditionOperand
	case "==":
		return registers[instr.conditionRegister] == instr.conditionOperand
	case "!=":
		return registers[instr.conditionRegister] != instr.conditionOperand
	}
	panic(fmt.Sprintf("Invalid condition [%s]", instr.condition))
}

func apply(registers map[string]int, inst instruction) {
	operand := inst.operand
	if inst.operation == "dec" {
		operand = -operand
	}
	registers[inst.register] = registers[inst.register] + operand
}

func parseInt(n string) int {
	i, err := strconv.ParseInt(n, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(i)
}

//go:embed input.txt
var input string

func main() {
	maximum, finalMaximum := solve(input)
	fmt.Printf("Final Maximum: %d, Total Maximum: %d\n", finalMaximum, maximum)
}
