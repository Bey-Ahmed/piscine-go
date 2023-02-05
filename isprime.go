package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	reste := 0
	stop := nb / 2
	for i := 2; i <= stop; i++ {
		reste = nb % i
		if reste == 0 {
			return false
		}
	}

	return true
}
