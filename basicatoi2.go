package piscine

func BasicAtoi2(s string) int {
	digits := []byte(s)
	size := len(digits)

	if size == 1 {
		if digits[0] >= '0' && digits[0] <= '9' {
			return int(digits[0] % '0')
		}
		return 0
	}

	intValue := 0
	position := 1
	for i := 0; i < size; i++ {
		if digits[i] < '0' || digits[i] > '9' {
			return 0
		}
		intValue = intValue + int(digits[size-1-i]%'0')*position
		position = position * 10
	}
	return intValue
}
