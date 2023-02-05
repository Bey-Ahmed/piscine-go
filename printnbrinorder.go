package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n > 0 {
		digits := []byte{}
		keep := n

		for keep != 0 {
			digits = append(digits, byte(keep%10))
			keep /= 10
		}

		sortByteSlice(digits)

		size := len(digits)
		for j := 0; j < size; j++ {
			z01.PrintRune(rune('0' + digits[j]))
		}
	}
}

func sortByteSlice(slice []byte) {
	size := len(slice)
	if size > 0 {
		lower := byte(0)
		for i := 0; i < size; i++ {
			lower = slice[i]
			for j := i + 1; j < size; j++ {
				if lower > slice[j] {
					lower = slice[j]
					slice[j] = slice[i]
					slice[i] = lower
				}
			}
		}
	}
}
