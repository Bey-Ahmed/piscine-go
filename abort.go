package piscine

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	SortIntegerTable(tab)
	size := len(tab)
	limit := size / 2
	sum := 0
	for i := 0; i < limit; i++ {
		if tab[i] >= 9223372036854775807-tab[i+1] || tab[i] <= -9223372036854775807-tab[i+1] {
			sum += 0
		} else {
			sum += tab[i] + tab[i+1]
		}
		if tab[size-2-i] >= 9223372036854775807-tab[size-1-i] || tab[size-2-i] <= -9223372036854775807-tab[size-1-i] {
			sum += 0
		} else {
			sum += tab[size-2-i] + tab[size-1-i]
		}
	}
	if size%2 != 0 {
		sum += tab[limit]
	}
	return sum
}
