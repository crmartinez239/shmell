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
	//	defer lex.Close()

	par := shm.NewParser(lex)
	pErr := par.Parse()
	if pErr != nil {
		switch e := pErr.(type) {

		case *shm.ParserError:
			fmt.Printf("Error: %d:%d - %s\n", e.Token().Line(), e.Token().Position(), e)
			return
		}
	}
	fmt.Println("BOOM!")
}
