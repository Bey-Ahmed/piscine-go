package piscine

func ForEach(f func(int), a []int) {
	for _, elmt := range a {
		f(elmt)
	}
}
