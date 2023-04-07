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

//Adding the value into the linked list (init operation)
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

//deleting data from the list
func (l *LinkedList) deleteNode(data int) {
	temp := l.head
	var prev *Node
	
	if temp != nil && temp.data == data {
		l.head = temp.next
		return
	}
	for temp != nil && temp.data != data {
		prev = temp
		temp = temp.next
	}
	if temp == nil {
		fmt.Println("Value not found")
		return
	}
	if temp == l.tail {
		l.tail = prev
		l.tail.next = nil
		return
	}
	prev.next = temp.next
}

//Insertion after a value
func (l *LinkedList) insertion(nextTo int, data int) {
	newNode := &Node{data, nil}
	temp := l.head

	for temp != nil && temp.data != nextTo {
		temp = temp.next
	}
	if temp == nil {
		return
	}
	if temp == l.tail {
		l.tail.next = newNode
		l.tail = newNode
		return
	}
	newNode.next = temp.next
	temp.next = newNode
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
	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(49)
	list.AddNode(33)

	list.printNode()
	fmt.Println("after deleting ...")
	// list.deleteNode(49)

	list.insertion(20, 30)
	list.printNode()
}
