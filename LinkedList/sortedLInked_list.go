package main

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

//Adding the value into the linked list (init operation)
func (l *LinkedList) AddNode(data int) {
	newNode := &ListNode{data, nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list3 *ListNode
	for list1.next != nil {
		if list1.data < list2.data {
			newList := &ListNode{list1.data, nil}
			list3 = newList
			list1 = list1.next
		} else {
			newList := &ListNode{list2.data, nil}
			list2 = list2.next
		}
	}
	return list3
}

func main() {
	list := LinkedList{}
	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(49)
	list.AddNode(33)

}
