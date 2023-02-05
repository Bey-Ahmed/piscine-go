package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argsNumb := len(os.Args)
	if argsNumb > 1 {
		start := 1
		init := 'a'
		if os.Args[1] == "--upper" {
			start = 2
			init = 'A'
		}
		for i := start; i < argsNumb; i++ {
			position := Atoi(os.Args[i])
			if position >= 1 && position <= 26 {
				z01.PrintRune(rune(int(init) + position - 1))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func Atoi(s string) int {
	digits := []byte(s)
	size := len(digits)

	if size == 0 {
		return 0
	}

	if size == 1 {
		if digits[0] >= '0' && digits[0] <= '9' {
			return int(digits[0] % '0')
		}
		return 0
	}

	intValue := 0
	position := 1
	start := 0
	if digits[0] == '+' || digits[0] == '-' {
		start = 1
	}

	for i := start; i < size; i++ {
		if digits[i] < '0' || digits[i] > '9' {
			return 0
		}

		if digits[0] == '-' {
			intValue = intValue - int(digits[size-1-i+start]%'0')*position
		} else {
			intValue = intValue + int(digits[size-1-i+start]%'0')*position
		}

		position = position * 10
	}
	return intValue
}
