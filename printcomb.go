package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for char1 := '0'; char1 <= '7'; char1++ {
		for char2 := char1 + 1; char2 <= '9'; char2++ {
			for char3 := char2 + 1; char3 <= '9'; char3++ {
				z01.PrintRune(char1)
				z01.PrintRune(char2)
				z01.PrintRune(char3)
				if char1 != '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
