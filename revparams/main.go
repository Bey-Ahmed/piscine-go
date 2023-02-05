package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argsNumb := len(os.Args)
	for i := argsNumb - 1; i > 0; i-- {
		argument := []rune(os.Args[i])
		argSize := len(argument)
		for i := 0; i < argSize; i++ {
			z01.PrintRune(argument[i])
		}
		z01.PrintRune('\n')
	}
}
