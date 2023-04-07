package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) push(data int) {
	newNode := &Node{data, nil}

	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
}

func (q *Stack) pop() {
	if q.top == nil {
		fmt.Println("emty")
	} else {
		q.top = q.top.next
	}
}

func (q *Stack) display() {
	curr := q.top

	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	q := Stack{}
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	q.push(5)
	q.push(6)
	q.display()
	q.pop()
	q.pop()
	q.display()

}
