package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		l.Tail = nil
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data_ref {
			if current.Next == l.Tail {
				l.Tail = current
			}
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}
