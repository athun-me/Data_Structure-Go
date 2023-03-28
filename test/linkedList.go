package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) addNode(data int) {
	newNode := &Node{data, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList) display() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func (l *LinkedList) insertAfter(data int, value int) {
	newNode := &Node{data, nil}
	curr := l.head

	for curr != nil && curr.data != value {
		curr = curr.next
	}
	if curr == nil {
		return
	}
	if curr == l.tail {
		l.tail.next = newNode
		l.tail = newNode
		return
	}
	newNode.next = curr.next
	curr.next = newNode
}

func (l *LinkedList) insertBefore(data int, value int) {
	newNode := &Node{data, nil}
	curr := l.head
	var prev *Node
	if l.head.data == value {
		newNode.next = l.head
		l.head = newNode
		return
	}
	for curr != nil && curr.data != value {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		return
	}
	newNode.next = curr
	prev.next = newNode
}

func (l *LinkedList) deleteNode(data int) {
	curr := l.head
	var prev *Node

	if curr != nil && l.head.data == data {
		l.head = curr.next
		return
	}
	for curr != nil && curr.data != data {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		return
	}
	if curr == l.tail {
		l.tail = prev
		l.tail.next = nil
		return
	}
	prev.next = curr.next
}

func (l *LinkedList) reverse() {
	curr := l.head
	var prev *Node

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}




func (l *LinkedList)ReverRec(){
	if l.head == nil{
		return
	}
	l.head = l.ReverRecHelper(l.head, l.head.next)
}

func (l *LinkedList)ReverRecHelper(node *Node, nextNode *Node)*Node{
	if nextNode == nil{
		// l.tail = node
		return node
	}
	newHead := l.ReverRecHelper(nextNode, nextNode.next)
	nextNode.next = node
	node.next = nil
	return newHead
}




func main() {
	list := LinkedList{}
	list.addNode(10)
	list.addNode(20)
	list.addNode(30)
	list.addNode(40)
	list.addNode(50)

	list.display()
	fmt.Println()
	list.ReverRec()
	list.display()
}
