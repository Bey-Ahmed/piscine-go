package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	size := len(a)
	for i := 0; i < size; i++ {
		length := len(a[i])
		for j := 0; j < length; j++ {
			z01.PrintRune(rune(a[i][j]))
		}
		z01.PrintRune('\n')
	}
}
