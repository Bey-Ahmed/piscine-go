package main

import "github.com/01-edu/z01"

func main() {
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	var alphabetLength int = len([]rune(alphabet))

	for i := 0; i < alphabetLength; i++ {
		z01.PrintRune(rune(alphabet[i]))
	}

	z01.PrintRune('\n')
}
