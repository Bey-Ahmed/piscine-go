package main

import (
	"os"
)

func main() {
	size := len(os.Args)

	if size != 4 || (os.Args[2] != "+" && os.Args[2] != "-" && os.Args[2] != "*" && os.Args[2] != "/" && os.Args[2] != "%") || !IsNumeric(os.Args[1]) || !IsNumeric(os.Args[3]) {
		return
	}

	if os.Args[2] == "/" && os.Args[3] == "0" {
		output := []byte("No division by 0\n")
		os.Stdout.Write(output)
		return
	}

	if os.Args[2] == "%" && os.Args[3] == "0" {
		output := []byte("No modulo by 0\n")
		os.Stdout.Write(output)
		return
	}

	operande1 := Atoi(os.Args[1])
	operande2 := Atoi(os.Args[3])
	if operande1 >= 9223372036854775807 || operande1 <= -9223372036854775807 || operande2 >= 9223372036854775807 || operande2 <= -9223372036854775807 {
		return
	}
	calcul := 0

	switch os.Args[2] {
	case "+":
		calcul = operande1 + operande2
	case "-":
		calcul = operande1 - operande2
	case "*":
		calcul = operande1 * operande2
	case "/":
		calcul = operande1 / operande2
	case "%":
		calcul = operande1 % operande2
	}

	printNbr(calcul)
	output := []byte("\n")
	os.Stdout.Write(output)
}

func printNbr(number int) {
	if number != 0 {
		store := []byte{}
		keep := number
		if number < 0 {
			output := []byte("-")
			os.Stdout.Write(output)
		}

		counter := 0
		for keep != 0 {
			if keep < 0 {
				store = append([]byte{byte('0' - (keep % 10))}, store...)
			} else {
				store = append([]byte{byte('0' + (keep % 10))}, store...)
			}
			counter++
			keep = keep / 10
		}

		os.Stdout.Write(store)
	} else {
		output := []byte("0")
		os.Stdout.Write(output)
	}
}

func IsNumeric(s string) bool {
	if len(s) <= 0 {
		return false
	}
	store := []rune(s)
	size := len(s)
	start := 0
	if store[0] == '-' {
		start = 1
	}
	for i := start; i < size; i++ {
		if store[i] < '0' || store[i] > '9' {
			return false
		}
	}
	return true
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
