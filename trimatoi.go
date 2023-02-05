package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	store := []byte(s)
	size := len(store)
	digits := []byte{}
	counter := 0
	for i := 0; i < size; i++ {
		if (counter == 0 && store[i] == '-') || (store[i] >= '0' && store[i] <= '9') {
			digits = append(digits, store[i])
			counter++
		}
	}

	return Atoi(string(digits))
}
