package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{
		Data: data,
		Next: nil,
	}

	if l.Tail != nil {
		l.Tail.Next = node
		l.Tail = node
	} else if l.Head == nil {
		l.Head = node
		l.Tail = l.Head
	} else {
		l.Tail = node
	}
}
