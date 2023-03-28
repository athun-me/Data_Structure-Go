package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) addnode(data int) {
	newNode := &Node{data, nil}
	if l.head == nil {
		l.head = newNode
	} else {
		curr := l.head
		for curr != nil {
			curr = curr.next
		}
		curr = newNode
	}
}
