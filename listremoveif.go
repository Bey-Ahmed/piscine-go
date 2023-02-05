package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		removeHead(l)
	}

	if l.Head == nil || l.Head.Next == nil {
		return
	}

	previous := l.Head
	courant := previous.Next
	for courant != nil {
		if courant.Data == data_ref {
			previous.Next = courant.Next
			courant.Next = nil
		} else {
			previous = courant
		}
		courant = previous.Next
	}
}

func removeHead(l *List) {
	courant := l.Head
	if l.Head == nil {
		return
	}
	l.Head = courant.Next
	if l.Head == l.Tail {
		l.Tail = nil
	}
	courant.Next = nil
}
