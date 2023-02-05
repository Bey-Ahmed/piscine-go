package piscine

func UltimateDivMod(a *int, b *int) {
	var keep int = *b
	*b = *a % *b
	*a = *a / keep
}
