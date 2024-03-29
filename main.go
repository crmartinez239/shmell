package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crmartinez239/shmell/shm"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	}

	lex, err := shm.NewLexer(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer lex.Close()

	par := shm.NewParser(lex)
	fmt.Printf("Error: %s", par.Parse().Error())
}
