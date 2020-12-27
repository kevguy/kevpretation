package main

import (
	"fmt"
	"github.com/kevguy/kevpretation/repl"
	"os"
)

func main() {
	fmt.Printf("Hi Auntie! This is the Kevpretation programming language!\n")
	fmt.Printf("Feel free to type in commands \n")
	repl.Start(os.Stdin, os.Stdout)
}
