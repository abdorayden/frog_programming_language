package main

import (
	"flag"
	"fmt"
	"os"

	"frog_programming_language/frog"
)

func main() {
	parse := flag.Bool("parse", false, "set to true to parse the file")

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: frog [options] <filepath>")
		flag.PrintDefaults()
		return
	}

	filepath := flag.Arg(0)

	code, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lexer := frog.NewLexer(string(code))

	if *parse {
		fmt.Println("Parsing the file...")
		parser := frog.NewParser(lexer)
		program := parser.ParseProgram()

		if parser.IsThereAnyErrors() {
			fmt.Println("Parser has errors:")
			for _, msg := range parser.Errors() {
				fmt.Println("\t" + msg)
			}
			return
		}

		fmt.Println("Generated AST:")
		fmt.Println(program.String())

	} else {
		tokens := lexer.GetAllTokens()
		for _, token := range tokens {
			fmt.Printf("%s: %q\n", frog.TokenToString(token.Type), token.Literal)
		}
	}
}
