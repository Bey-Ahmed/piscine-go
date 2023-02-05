package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}

	if nb == 1 {
		return 1
	}

	result := nb / 2
	for result != 0 && result*result != nb {
		approx := (result + (nb / result)) / 2
		if (result*result < nb && approx*approx < nb) || (result*result < nb && approx*approx > nb) {
			return 0
		}
		result = approx
	}
	return result
}
