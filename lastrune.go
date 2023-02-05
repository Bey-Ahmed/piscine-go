package piscine

/*
* Retourne le dernier caractère d'une chaîne
 */
func LastRune(s string) rune {
	// Récupération et stockage de la chaîne dans un tableau de runes
	store := []rune(s)

	// Retour de la dernière valeur donc du dernier caractère de ce rune
	return store[len(store)-1]
}
