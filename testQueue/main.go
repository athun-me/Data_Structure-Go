package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) enqueue(data int) {
	newNode := &Node{data, nil}
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *Queue) dequeue() {
	if q.front == nil {
		fmt.Println("list empty")
	} else {
		q.front = q.front.next
	}
}

func (q *Queue) display() {
	curr := q.front
	if q.front == nil {
		fmt.Println("emty")
		return
	}
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	q := Queue{}
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.enqueue(40)
	q.enqueue(50)
	q.display()
	q.dequeue()
	q.display()
}
