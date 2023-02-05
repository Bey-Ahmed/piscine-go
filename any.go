package piscine

func Any(f func(string) bool, a []string) bool {
	for _, elmt := range a {
		if f(elmt) {
			return true
		}
	}
	return false
}
