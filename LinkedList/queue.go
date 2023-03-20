package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type queue struct {
	front *Node
	rear  *Node
}

func (l *queue) enqueue(data int) {
	newNode := Node{data, nil}

	if l.rear == nil {
		l.front = &newNode
		l.rear = &newNode
	} else {
		l.rear.next = &newNode
		l.rear = &newNode
	}
}

func (l *queue) denqueue() {
	if l.front == nil {
		fmt.Println("empty")
	} else {
		l.front = l.front.next
	}

	if l.front == nil {
		l.rear = nil
		return
	}
}

func (l *queue) display() {
	current := l.front

	if l.front == nil {
		fmt.Println("list is empty")
		return
	} else {
		for current != nil {
			fmt.Printf("%d\t", current.data)
			current = current.next
		}
	}
	fmt.Println()
}

func main() {
	q := queue{}
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.enqueue(40)
	q.enqueue(50)

	q.display()

	q.denqueue()
	q.display()
	q.denqueue()
	q.display()

}
