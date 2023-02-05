package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	size := len(base)
	if size == 0 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	digits := []rune{}
	keep := nbr
	if keep == 0 {
		digits = append(digits, rune(0))
	}
	for keep != 0 {
		value := 0
		if keep > 0 {
			value = keep % size
		} else if keep < 0 {
			value = -keep % size
		}
		digits = append(digits, rune(base[value]))
		keep /= size
	}
	digitsNbr := len(digits)
	if nbr < 0 {
		z01.PrintRune('-')
	}
	for k := digitsNbr - 1; k >= 0; k-- {
		z01.PrintRune(digits[k])
	}
}
