package piscine

/*
* Compare 2 chaînes données en arguments
*/
func Compare(a, b string) int {
	// Retourne -1 si la première chaîne est plus petite
	if a < b {
		return -1
	// Retourne 1 si la première chaîne est plus grande
	} else if a > b {
		return 1
	// Retourne 0 si elles sont égales
	} else {
		return 0
	}
}
