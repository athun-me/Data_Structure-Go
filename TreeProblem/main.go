package main

import (
	"fmt"
	"math"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}
type TreeNode struct {
	root *Node
}

func (t *TreeNode) insert(val int) {
	newNode := &Node{val: val}
	if t.root == nil {
		t.root = newNode
		return
	}
	t.root.insert(newNode)

}
func (n *Node) insert(newNode *Node) {
	if n.val > newNode.val {
		if n.left == nil {
			n.left = newNode
			return
		} else {
			n.left.insert(newNode)
		}
	} else {
		if n.val < newNode.val {
			if n.right == nil {
				n.right = newNode
				return
			} else {
				n.right.insert(newNode)
			}
		}
	}
}
func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Println(node)
		inOrder(node.right)
	}
}

func depth(root *Node) int {
	if root != nil {
		left := depth(root.left)
		right := depth(root.right)
		if left < right {
			return right + 1
		} else {
			return left + 1
		}
	}
	return 0
}
func isSymmetri(root *Node) bool {
	if root == nil {
		return true
	}
	return isMirror(root.left, root.right)
}

func isMirror(left, right *Node) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return (left.val == right.val) &&
		isMirror(left.left, right.right) &&
		isMirror(left.right, right.left)
}

func isSame(root1, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return (root1.val == root2.val) &&
		(isSame(root1.left, root2.left)) &&
		(isSame(root1.right, root2.right))
}

func BFS(root *Node) {
	if root == nil {
		return
	}
	que := []*Node{root}

	for len(que) > 0 {
		for i := 0; i < len(que); i++ {
			node := que[0]
			que = que[1:]
			fmt.Println(node.val)
			if node.left != nil {
				que = append(que, node.left)
			}
			if node.right != nil {
				que = append(que, node.right)
			}
		}
	}
}

func isValidBST(root *Node) bool {
	var def func(root *Node, min, max int) bool

	def = func(root *Node, min, max int) bool {
		if root == nil {
			return true
		}
		if root.val <= min || root.val >= max {
			return false
		}
		return def(root.left, min, root.val) && def(root.right, root.val, max)
	}

	return def(root, math.MinInt, math.MaxInt)
}

func preOrder(root *Node) {
	if root != nil {
		preOrder(root.left)
		preOrder(root.right)
		fmt.Println(root.val)
	}
}

func kthsmallets(root *Node, k int) int {
	var res int
	count := 0
	inOrderTraversal(root, &count, &k, &res)
	return res
}
func inOrderTraversal(root *Node, count, k, result *int) {

	if root == nil {
		return
	}

	inOrderTraversal(root.left, count, k, result)
	*count++
	if *k == *count {
		*result = root.val
		return
	}
	inOrderTraversal(root.right, count, k, result)

}

func deletion(root *Node, val int) *Node {
	if root == nil {
		return nil
	}

	if root.val > val {
		root.left = deletion(root.left, val)
	} else if root.val < val {
		root.right = deletion(root.right, val)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		root.val = findMin(root.right)
		root.right = deletion(root.right, root.val)
	}
	return root
}

func findMin(node *Node) int {
	for node.left != nil {
		node = node.left
	}
	return node.val
}

func main() {

	tree := &TreeNode{}
	tree.insert(4)
	tree.insert(2)
	tree.insert(7)
	tree.insert(1)
	tree.insert(3)
	tree.insert(9)
	tree.insert(6)
	// fmt.Println(kthsmallets(tree.root, 3))

	inOrder(deletion(tree.root, 4))

}
