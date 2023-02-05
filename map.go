package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, elmt := range a {
		result[i] = f(elmt)
	}
	return result
}
