package piscine

func RecursiveFactorial(number int) int {
	if number == 0 {
		return 1
	}

	if number < 0 || number > 65536 {
		return 0
	}

	return number * RecursiveFactorial(number-1)
}
