package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for char1 := '0'; char1 <= '9'; char1++ {
		for char2 := '0'; char2 <= '9'; char2++ {
			for char3 := char1; char3 <= '9'; char3++ {
				for char4 := '0'; char4 <= '9'; char4++ {
					value1 := char1 % '0'
					value2 := char2 % '0'
					value3 := char3 % '0'
					value4 := char4 % '0'

					store1 := value1*10 + value2
					store2 := value3*10 + value4

					if store1 < store2 {
						z01.PrintRune(char1)
						z01.PrintRune(char2)
						z01.PrintRune(' ')
						z01.PrintRune(char3)
						z01.PrintRune(char4)
						if char1 != '9' || char2 != '8' || char3 != '9' || char4 != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
