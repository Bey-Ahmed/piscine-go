package piscine

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}

	if l.Head == l.Tail {
		return 1
	}

	courant := l.Head
	counter := 1
	for courant.Next != nil {
		counter++
		courant = courant.Next
	}

	return counter
}
