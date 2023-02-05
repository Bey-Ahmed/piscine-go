package piscine

func ActiveBits(n int) int {
	if n == 0 {
		return 0
	}

	counter := 0
	for n != 0 {
		if n%2 == 1 {
			counter++
		}
		n = n / 2
	}

	return counter
}
