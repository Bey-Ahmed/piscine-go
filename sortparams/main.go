package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argsNumb := len(os.Args)
	arguments := []string(os.Args)
	for i := 1; i < argsNumb-1; i++ {
		for j := i + 1; j < argsNumb; j++ {
			if arguments[i] > arguments[j] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}

	for n := 1; n < argsNumb; n++ {
		argument := []rune(os.Args[n])
		argSize := len(argument)
		for p := 0; p < argSize; p++ {
			z01.PrintRune(argument[p])
		}
		z01.PrintRune('\n')
	}
}
