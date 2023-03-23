package main

import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node 
	tail *Node
}

func (l *LinkedList)addNode(data int){
	newNode := &Node{data, nil}
	
	if l.head == nil{
		l.head = newNode
		l.tail = newNode
		return
	}else{
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList)displayList(){
	curr := l.head
	for curr != nil{
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func main(){
	var array = [10]int{1,20,43,21,54,10,34,98,42,65}
	list := LinkedList{}

	for i:=0; i<len(array); i++{
		list.addNode(array[i])
	}
	
	list.displayList()
}