package piscine

func Swap(a *int, b *int) {
	var keep int = *a
	*a = *b
	*b = keep
}
