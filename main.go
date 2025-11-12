package main

import (
	"fmt"
	// "log"
	"flag"
	"os"
)

func main() {
	lex := flag.Bool("lex", false, "set to true to tokenize the file")
	parse := flag.Bool("parse", false, "set to true to parse the file")

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: frog [options] <filepath>")
		flag.PrintDefaults()
		return
	}

	filepath := flag.Arg(0)


	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))

	if *lex {
		fmt.Println("Lexing the file...")
	}

	if *parse {
		fmt.Println("Parsing the file...")
	}
}
