package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos == 0 {
		return l
	}

	courant := l
	counter := 0
	for courant != nil {
		if counter == pos {
			return courant
		}
		counter++
		courant = courant.Next
	}

	return nil
}
