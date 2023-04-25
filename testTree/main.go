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
		
	}
}

func (t *tree) delete(data int) {
	t.deleteHelper(data, t.root, nil)
}
func (t *tree) deleteHelper(data int, currentNode *Node, parentNode *Node) {
	for currentNode != nil {
		if data < currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if data > currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.right
		} else {
			if currentNode.left != nil && currentNode.right != nil {
				currentNode.data = t.getMin(currentNode.right)
				t.deleteHelper(currentNode.data, currentNode.right, currentNode)
			} else {
				if parentNode == nil {
					if currentNode.right == nil {
						t.root = currentNode.left
					} else {
						t.root = currentNode.right
					}
				} else {
					if parentNode.left == currentNode {
						if currentNode.right == nil {
							parentNode.left = currentNode.left
						} else {
							parentNode.left = currentNode.right
						}
					} else {
						if currentNode.right == nil {
							parentNode.right = currentNode.left
						} else {
							parentNode.right = currentNode.right
						}
					}
				}
			}
			break
		}
	}
}
func (t *tree) getMin(Node *Node) int {
	if Node.left == nil {
		return Node.data
	} else {
		return t.getMin(Node.left)
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

func (t *tree) preOder() {
	t.preOderHelper(t.root)
	fmt.Println()
}

func (t *tree) preOderHelper(Node *Node) {
	if Node == nil {
		return
	}
	fmt.Printf("%d ", Node.data)
	t.preOderHelper(Node.left)
	t.preOderHelper(Node.right)
}

func (t *tree) postOder() {
	t.postOderHelper(t.root)
	fmt.Println()
}
func (t *tree) postOderHelper(Node *Node) {
	if Node == nil {
		return
	}
	t.postOderHelper(Node.left)
	t.postOderHelper(Node.right)
	fmt.Printf("%d ", Node.data)
}

func (t *tree) findClosest(target int) int {
	currentNode := t.root
	closest := currentNode.data
	for currentNode != nil {
		if convert(target-closest) > convert(target-currentNode.data) {
			closest = currentNode.data
		}
		if target < currentNode.data {
			currentNode = currentNode.left
		} else if target > currentNode.data {
			currentNode = currentNode.right
		} else {
			break
		}
	}
	return closest

}

func convert(data int) int {
	var newData int
	if data < 0 {
		newData = data * -1
	} else {
		return data
	}
	return newData
}

func main() {
	t := tree{}
	t.insert(60)
	t.insert(70)
	t.insert(20)
	t.insert(50)
	t.insert(10)

	// t.inOder()
	// fmt.Println()
	// t.preOder()
	// fmt.Println()
	// t.postOder()
	fmt.Println(t.findClosest(60))
}
