package main

import "fmt"

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

func (s *Stack) pop() {
	if s.top == nil {
		fmt.Println("list emty")
	} else {
		s.top = s.top.next
	}
}

func (s *Stack) display() {
	curr := s.top
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	s := Stack{}
	s.push(10)
	s.push(20)
	s.push(30)
	s.push(40)
	s.display()
	s.pop()
	s.display()
}
