package main

import (
	"log"
	"os"

	"github.com/crmartinez239/lang1/lang"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	}

	lex, err := lang.NewLexer(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer lex.Close()

}
