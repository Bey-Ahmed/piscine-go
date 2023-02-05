package piscine

func ListReverse(l *List) {
	reverseL := &List{
		Head: nil,
		Tail: nil,
	}

	if l.Head == nil {
		return
	}

	reverseL.Tail = &NodeL{
		Data: l.Head.Data,
		Next: nil,
	}

	reverseL.Head = reverseL.Tail

	courant := l.Head.Next
	keep := reverseL.Head
	for courant != nil {
		reverseL.Head = &NodeL{
			Data: courant.Data,
			Next: keep,
		}
		keep = reverseL.Head
		courant = courant.Next
	}

	*l = *reverseL
}
