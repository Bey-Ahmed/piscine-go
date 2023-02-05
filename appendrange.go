package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	intSlice := []int{}
	size := max - min
	for i := 0; i < size; i++ {
		intSlice = append(intSlice, min+i)
	}

	return intSlice
}
