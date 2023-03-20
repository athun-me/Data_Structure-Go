package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type stack struct {
	top *Node
}

func (l *stack) push(data int) {
	newNode := Node{data, nil}
	if l.top == nil {
		l.top = &newNode
	} else {
		newNode.next = l.top
		l.top = &newNode
	}
}

func (l *stack) display() {
	current := l.top

	for current != nil {
		fmt.Printf("%d\n", current.data)
		current = current.next
	}
}

func (l *stack) pop() {
	if l.top == nil {
		fmt.Println("stack underflow")
		return
	}
	l.top = l.top.next
}

func main() {
	stack := stack{}

	stack.push(10)
	stack.push(20)
	stack.push(30)
	stack.push(40)
	stack.push(50)

	stack.display()
	stack.pop()
	fmt.Printf("\n")
	stack.display()
}
