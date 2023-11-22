package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type ListNode struct {
	head *Node
}

func (l *ListNode) insert(val int) {
	newNode := &Node{Val: val, Next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}
func (l *ListNode) deleteNode(val int) {
	if l.head.Val == val {
		l.head = l.head.Next
		return
	}

	cur := l.head

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
	}

}

func (l *ListNode) insertAfter(val int, data int) {
	newNode := &Node{Val: data, Next: nil}

	if l.head.Val == val {
		newNode.Next = l.head.Next
		l.head.Next = newNode
		return
	}

	cur := l.head

	for cur.Next != nil {
		if cur.Val == val {
			newNode.Next = cur.Next
			cur.Next = newNode
			return
		}
		cur = cur.Next
	}
	cur.Next = newNode
}

func (l *ListNode) insertAtBeginningAtaValu(val int, data int) {
	newNode := &Node{Val: data, Next: nil}
	if l.head.Val == val {
		newNode.Next = l.head
		l.head = newNode
		return
	}
	cur := l.head
	for cur.Next != nil {
		if cur.Next.Val == val {
			newNode.Next = cur.Next
			cur.Next = newNode
			return
		}
		cur = cur.Next
	}

}

func (l *ListNode) insertAtBeginning(val int) {
	newNode := &Node{Val: val, Next: nil}
	if l.head != nil {
		newNode.Next = l.head
		l.head = newNode
	}
}

func (l *ListNode) display() {
	cur := l.head

	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func (l *ListNode) swap(val int) {
	if l.head.Val == val {
		temp := l.head.Next
		l.head.Next = l.head.Next.Next
		temp.Next = l.head
		l.head = temp
		return
	}

	cur := l.head

	for cur.Next.Next != nil {
		if cur.Next.Val == val {
			temp := cur.Next.Next
			cur.Next.Next = temp.Next
			temp.Next = cur.Next
			cur.Next = temp
			return
		}
		cur = cur.Next
	}
}

func merge(list1 *Node, list2 *Node) *Node {
	res := &Node{}
	curr := res

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = &Node{Val: list1.Val, Next: nil}
			list1 = list1.Next
		} else {
			curr.Next = &Node{Val: list2.Val, Next: nil}
			list2 = list2.Next
		}

		curr = curr.Next
	}
	if list1 != nil {
		curr = list1
	}

	if list2 != nil {
		curr = list2
	}

	return res.Next
}
func (l *ListNode) revers() {
	curr := l.head
	var prev *Node = nil

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	l.head = prev
}

func (l *ListNode) recursiveReverse() {
	l.head = recursiveReverse(l.head, nil)
}

func recursiveReverse(curr *Node, prev *Node) *Node {
	if curr == nil {
		return prev
	}

	next := curr.Next
	curr.Next = prev
	return recursiveReverse(next, curr)
}

func concatenateLists(lists []*Node) *Node {
	head := lists[0]
	var list *Node
	var end *Node
	for _, list = range lists {
		if end != nil {
			end.Next = list
			end = nil
		}
		for list.Next != nil {
			list = list.Next
		}
		if list.Next == nil {
			end = list
		}

	}

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	return nil
}

func removeNthFromEnd(head *Node, n int) *Node {
	temp := head
	length := 0
	for temp != nil {
		length++
		temp = temp.Next
	}

	if length == n {
		head = head.Next
		return head
	}

	i := 1
	cur := head
	for i != length-n {
		i++
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return head
}
func x(arr1 []int, arr2 []int) []int {
	res := []int{}
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	if i != len(arr1)-1 {
		res = append(res, arr1[i:]...)
	}
	if j != len(arr2)-1 {
		res = append(res, arr2[j:]...)
	}

	return res

}

func main() {
	list := ListNode{}
	list.insert(10)
	list.insert(20)
	list.insert(30)
	list.insert(40)
	list.insert(50)

	// list.insertAtBeginning(200)
	// list.deleteNode(50)
	// list.insertAtBeginningAtaValu(50, 2000)
	// list.insertAfter(50, 200)
	list.swap(50)
	list.display()

}
