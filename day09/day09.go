package main

import (
	_ "embed"
	"fmt"
)

func solve(input string) lexer {

	lex := lexer{source: input}

	for lex.index < len(lex.source) {
		switch lex.source[lex.index] {
		case '<':
			lex.skipGarbage()
		case '{':
			lex.openGroup()
		case '}':
			lex.closeGroup()
		case ',':
			lex.index++
		default:
			panic(fmt.Sprintf("Invalid character: [%c]\n", lex.source[lex.index]))
		}
	}

	return lex
}

type lexer struct {
	source       string
	index        int
	currentScore int
	totalScore   int
	garbage      int
}

func (lex *lexer) skipGarbage() {
	lex.index++
	for lex.index < len(lex.source) && lex.source[lex.index] != '>' {
		if lex.source[lex.index] == '!' {
			lex.index += 2
		} else {
			lex.index++
			lex.garbage++
		}
	}
	lex.index++
}

func (lex *lexer) openGroup() {
	lex.currentScore++
	lex.index++
}

func (lex *lexer) closeGroup() {
	lex.totalScore += lex.currentScore
	lex.currentScore--
	lex.index++
}

//go:embed input.txt
var input string

func main() {
	lex := solve(input)
	fmt.Printf("Score: %d, Garbage: %d\n", lex.totalScore, lex.garbage)
}
