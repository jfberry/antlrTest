package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"time"
)
import "antlrTest/pokemonSearchParser"

func main() {
	// Setup the input
	is := antlr.NewInputStream("(a0|a15)&d15&s14-15,0&l1")

	lexer := pokemonSearchParser.NewAdvancedSearchLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := pokemonSearchParser.NewAdvancedSearchParser(stream)
	parser.BuildParseTrees = true
	tree := parser.Entry()

	visitor := PokemonEvalVisitor{
		pokemonData: &PokemonData{
			Iv:    1,
			Atk:   1,
			Def:   15,
			Sta:   15,
			Level: 1,
		},
	}
	var result bool

	start := time.Now()
	for x := 0; x < 100000; x++ {
		result = visitor.Visit(tree)
	}

	fmt.Printf("res = %t time = %s", result, time.Since(start))
}
