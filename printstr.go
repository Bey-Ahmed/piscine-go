package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	runeArray := []rune(s)
	size := len(runeArray)

	for i := 0; i < size; i++ {
		z01.PrintRune(runeArray[i])
	}
}
