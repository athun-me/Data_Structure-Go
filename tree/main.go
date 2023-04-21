package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type tree struct {
	root *Node
}

func (t *tree) insert(data int) {
	newNode := &Node{data, nil, nil}
	currentNode := t.root
	if currentNode == nil {
		t.root = newNode
		return
	}
	for true {
		if data < currentNode.data {
			if currentNode.left == nil {
				currentNode.left = newNode
				break
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				currentNode.right = newNode
				break
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

func (t *tree) contains(data int) bool {
	currentNode := t.root
	for currentNode != nil {
		if data < currentNode.data {
			currentNode = currentNode.left
		} else if data > currentNode.data {
			currentNode = currentNode.right
		} else {
			return true
		}
	}
	return false
}

func (t *tree) delete(data int) {
	t.deleteHelper(data, t.root, nil)
}

func (t *tree) deleteHelper(data int, currendNode *Node, parentNode *Node) {
	for currendNode != nil {
		if data < currendNode.data {
			parentNode = currendNode
			currendNode = currendNode.left
		} else if data > currendNode.data {
			parentNode = currendNode
			currendNode = currendNode.right
		} else {
			if currendNode.left != nil && currendNode.right != nil {
				currendNode.data = t.getMin(currendNode.right)
				t.deleteHelper(currendNode.data, currendNode.right, currendNode)
			} else {
				if parentNode == nil {
					if currendNode.right == nil {
						t.root = currendNode.left
					} else {
						t.root = currendNode.right
					}
				} else {
					if parentNode.left == currendNode {
						if currendNode.right == nil {
							parentNode.left = currendNode.left
						} else {
							parentNode.left = currendNode.right
						}
					} else {
						if currendNode.right == nil {
							parentNode.right = currendNode.left
						} else {
							parentNode.right = currendNode.right
						}
					}
				}
			}
			break
		}
	}
}

func (t *tree) getMin(currentNode *Node) int {
	if currentNode.left == nil {
		return currentNode.data
	} else {
		return t.getMin(currentNode.left)
	}
}

func (t *tree) inOder() {
	t.inOderHelper(t.root)
	fmt.Println()
}
func (t *tree) inOderHelper(Node *Node) {
	if Node == nil {
		return
	}
	t.inOderHelper(Node.left)
	fmt.Printf("%d ", Node.data)
	t.inOderHelper(Node.right)
}

func main() {
	t := tree{}

	t.insert(11)
	t.insert(20)
	t.insert(10)
	t.insert(5)
	t.insert(25)
	t.insert(21)
	t.insert(3)

	t.inOder()

	t.delete(20)

	t.inOder()
	// fmt.Println(t.contains(3))

}
