package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(val int) {
	if t.root == nil {
		t.root = &Node{Val: val, Left: nil, Right: nil}
	} else {
		insert(t.root, val)
	}
}

func insert(root *Node, val int) {
	if root.Val < val {
		if root.Right == nil {
			root.Right = &Node{Val: val, Left: nil, Right: nil}
		} else {
			insert(root.Right, val)
		}
	} else {
		if root.Left == nil {
			root.Left = &Node{Val: val, Left: nil, Right: nil}
		} else {
			insert(root.Left, val)
		}
	}
}

func (t *Tree) inOrderTravers(root *Node) {
	if root != nil {
		t.inOrderTravers(root.Left)
		fmt.Println(root.Val)
		t.inOrderTravers(root.Right)
	}
}

func (t *Tree) preeOrder(root *Node) {
	if root != nil {
		fmt.Println(root.Val)
		t.preeOrder(root.Left)
		t.preeOrder(root.Right)
	}
}

func (t *Tree) PostOrder(root *Node) {
	if root != nil {
		t.PostOrder(root.Left)
		t.PostOrder(root.Right)
		fmt.Println(root.Val)
	}
}

func main() {
	arr := []int{4, 2, 3, 5, 6, 9, 8, 7, 1}
	t := &Tree{}
	for _, val := range arr {
		t.insert(val)
	}
	t.PostOrder(t.root)
}
