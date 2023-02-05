package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, elmt := range tab {
		if f(elmt) {
			counter++
		}
	}
	return counter
}
