package piscine

func IterativeFactorial(number int) int {
	if number < 0 {
		return 0
	}

	if number == 0 {
		return 1
	}

	factorial := 1
	for i := 2; i <= number; i++ {
		factorial *= i
		if factorial == 0 {
			return 0
		}
	}

	if factorial != int(factorial) {
		return 2
	}

	return factorial
}
