package piscine

func SortIntegerTable(table []int) {
	size := len(table)

	for i := 0; i < size-1; i++ {
		keep := table[i]
		for j := i + 1; j < size; j++ {
			if keep > table[j] {
				keep = table[j]
				table[j] = table[i]
				table[i] = keep
			}
		}
	}
}
