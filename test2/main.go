package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedLists struct {
	head *Node
	tail *Node
}

func (l *LinkedLists) addNode(data int) {
	newNode := &Node{data, nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedLists) display() {
	curr := l.head

	for curr != nil {
		fmt.Println(curr.data, curr.next)
		curr = curr.next
	}
}

func main() {
	list := LinkedLists{}
	var array = [10]int{1, 29, 39, 43, 55, 65, 78, 89, 90, 122}
	for i := 0; i < len(array)-1; i++ {
		list.addNode(array[i])
	}
	list.display()
}
