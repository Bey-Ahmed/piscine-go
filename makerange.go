package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min
	intSlice := make([]int, size)
	for i := 0; i < size; i++ {
		intSlice[i] = min + i
	}

	return intSlice
}
