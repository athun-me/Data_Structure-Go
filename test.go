package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(data int) {
	newNode := &Node{data, nil, nil}
	currentNode := t.root
	if currentNode == nil {
		t.root = newNode
		return
	}
	for currentNode != nil {
		if data < currentNode.data {
			if currentNode.left == nil {
				currentNode.left = newNode
				return
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				currentNode.right = newNode
				return
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

func (t *Tree) inOder() {
	t.inOderHelper(t.root)
	fmt.Println()
}
func (t *Tree) inOderHelper(currentNode *Node) {
	if currentNode == nil {
		return
	}
	t.inOderHelper(currentNode.left)
	fmt.Printf("%d ", currentNode.data)
	t.inOderHelper(currentNode.right)
}

func (t *Tree) remove(data int) {
	t.removeHelper(data, t.root, nil)
}
func (t *Tree) removeHelper(data int, currentNode *Node, parentNode *Node) {
	for currentNode != nil {
		if data < currentNode.data {
			currentNode = currentNode.left
		} else if data > currentNode.data {
			currentNode = currentNode.right
		} else {
			if currentNode.left != nil && currentNode.right != nil {
				currentNode.data = t.getMin(currentNode.right)
				t.removeHelper(currentNode.data, currentNode.right, currentNode)
			} else {
				if parentNode == nil {
					if currentNode.left == nil {
						t.root = currentNode.right
					} else {
						t.root = currentNode.left
					}
				} else {
					if parentNode.left == currentNode {
						if currentNode.left == nil {
							parentNode.left = currentNode.right
						} else {
							parentNode.left = currentNode.left
						}
					} else {
						if currentNode.left == nil {
							parentNode.right = currentNode.right
						} else {
							parentNode.right = currentNode.left
						}
					}
				}
			}
			break
		}
	}
}

func (t *Tree) getMin(currentNode *Node) int {
	if currentNode.left == nil {
		return currentNode.data
	} else {
		return t.getMin(currentNode.left)
	}
}

func main() {
	tree := Tree{}

	tree.insert(11)
	tree.insert(20)
	tree.insert(10)
	tree.insert(5)
	tree.insert(25)
	tree.insert(21)
	tree.insert(3)

	tree.inOder()

	tree.remove(20)

	tree.inOder()

}
