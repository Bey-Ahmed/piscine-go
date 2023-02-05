package piscine

import "github.com/01-edu/z01"

func PrintNbr(number int) {
	if number != 0 {
		store := []rune{}
		keep := number
		if number < 0 {
			z01.PrintRune('-')
		}

		counter := 0
		for keep != 0 {
			if keep < 0 {
				store = append(store, rune('0'-(keep%10)))
			} else {
				store = append(store, rune('0'+(keep%10)))
			}
			counter++
			keep = keep / 10
		}

		size := len(store)
		for i := size - 1; i >= 0; i-- {
			z01.PrintRune(rune(store[i]))
		}
	} else {
		z01.PrintRune('0')
	}
}
