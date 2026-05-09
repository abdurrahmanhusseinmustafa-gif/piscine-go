package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	// 1. Create the new node
	newNode := &NodeL{Data: data}

	// 2. If the list is empty, this node is both Head and Tail
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	// 3. Point new node to the current head
	newNode.Next = l.Head

	// 4. Update the list head to be the new node
	l.Head = newNode
}
