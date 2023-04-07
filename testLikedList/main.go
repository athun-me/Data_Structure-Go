package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddNode(data int) {
	newNode := &Node{data, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

//displaying the elements in the list
func (l *LinkedList) printNode() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d\n", currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *LinkedList) deleteGivenValue(data int) {
	
	for l.head != nil && l.head.data == data {
		l.head = l.head.next
	}

	curr := l.head

	for curr.next != nil {
		if curr.next.data == data {
			curr.next = curr.next.next
		}else{
			curr = curr.next
		}
	}
}


func main() {
	list := LinkedList{}
	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(10)
	list.AddNode(10)
	list.AddNode(10)

	list.AddNode(30)
	list.AddNode(10)
	list.AddNode(15)
	list.AddNode(10)

	list.AddNode(10)

	list.printNode()

	list.deleteGivenValue(10)
	fmt.Println()
	list.printNode()

}
