package main

import (
	"flag"
	"fmt"
	"os"

	// for signals
	"os/signal"
	"syscall"

	"frog_programming_language/frog"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nCtrl+C pressed. Exiting.")
		os.Exit(0)
	}()

	parse := flag.Bool("parse", false, "set to true to parse the file")
	lex := flag.Bool("lex", false, "set to true to lex the file")
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

	if *lex {
		tokens := lexer.GetAllTokens()
		for _, token := range tokens {
			fmt.Printf("%s: %q\n", frog.TokenToString(token.Type), token.Literal)
		}
	} else if *parse {
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
		frog.PrintAST(program, "", true)

	} else {
		parser := frog.NewParser(lexer)
		program := parser.ParseProgram()

		if parser.IsThereAnyErrors() {
			fmt.Println("Parser has errors:")
			for _, msg := range parser.Errors() {
				fmt.Println("\t" + msg)
			}
			return
		}

		env := frog.NewEnvironment()
		evaluated := frog.Eval(program, env)
		if evaluated != nil {
			fmt.Println(evaluated.Inspect())
		}
	}
}
