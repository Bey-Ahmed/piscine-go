package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := []rune(os.Args[0])
	size := len(arguments)
	for i := 2; i < size; i++ {
		z01.PrintRune(arguments[i])
	}
	z01.PrintRune('\n')
}
