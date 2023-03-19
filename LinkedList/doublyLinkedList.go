package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

//creating new doubly linked list (init)
func (l *LinkedList) addNode(data int) {
	newNode := &Node{data, nil, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

//displaying
func (l *LinkedList) display() {
	temp := l.head
	if l.head == nil {
		return
	}
	for temp != nil {
		fmt.Printf("%d\n", temp.data)
		temp = temp.next
	}
}

//displaying in revers order
func (l *LinkedList) displayRevers() {
	temp := l.tail

	if l.tail == nil {
		return
	}
	for temp != nil {
		fmt.Printf("%d\n", temp.data)
		temp = temp.prev
	}
}

//delete Node
func (l *LinkedList) delete(data int) {
	temp := l.head
	after := temp.next
	var prev *Node
	if temp != nil && temp.data == data {
		l.head = temp.next
		return
	}
	for temp != nil && temp.data != data {
		after = after.next
		prev = temp
		temp = temp.next
	}

	if temp == l.tail {
		l.tail = temp.prev
		l.tail.next = nil
		return
	}
	prev.next = temp.next
	after.prev = temp.prev
}

//Insert after
func (l *LinkedList) insertAfter(nextTo int, data int) {
	newNode := &Node{data, nil, nil}
	temp := l.head
	after := temp.next

	for temp != nil && temp.data != nextTo {
		after = after.next
		temp = temp.next
	}
	if temp == nil {
		return
	}
	if temp == l.tail {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		return
	}
	newNode.next = after
	newNode.prev = temp
	temp.next = newNode
	after.prev = newNode
}

func (l *LinkedList) removeAdjesentDuplicates() {
	curr := l.head

	for curr != nil {
		next := curr.next
		for next != nil && next.data == curr.data {
			next = next.next
		}
		curr.next = next

		if next == l.tail && next.data == curr.data{
			l.tail = curr
			l.tail = nil
		}
		curr = next
	}
}

func main() {
	list := LinkedList{}

	list.addNode(10)
	list.addNode(90)
	list.addNode(30)
	list.addNode(60)
	list.addNode(60)
	list.addNode(60)
	list.addNode(50)
	list.addNode(50)

	list.display()

	fmt.Printf("\n")
	list.removeAdjesentDuplicates()
	list.display()
}
