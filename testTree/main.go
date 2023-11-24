package main

import "fmt"

type Node struct {
	Val   int
	left  *Node
	right *Node
}
type Tree struct {
	root *Node
}

func (t *Tree) insert(val int) {
	newNode := &Node{Val: val, left: nil, right: nil}
	if t.root == nil {
		t.root = newNode

	} else {
		insertHelper(t.root, val)
	}

}
func insertHelper(root *Node, val int) {
	newNoe := &Node{Val: val, left: nil, right: nil}
	if val < root.Val {
		if root.left == nil {
			root.left = newNoe
		} else {
			insertHelper(root.left, val)
		}
	} else {
		if root.right == nil {
			root.right = newNoe
		} else {
			insertHelper(root.right, val)
		}
	}
}
func (t *Tree) inOrderTravers(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}
	return (root1.Val == root2.Val) && t.inOrderTravers(root1.left, root2.left) && t.inOrderTravers(root1.right, root2.right)

}

func main() {
	tree := Tree{}
	values := []int{1, 2, 1}
	for _, value := range values {
		tree.insert(value)
	}
	tree2 := Tree{}
	values2 := []int{1, 1, 2}
	for _, value := range values2 {
		tree2.insert(value)
	}

	if tree.inOrderTravers(tree.root, tree2.root) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
