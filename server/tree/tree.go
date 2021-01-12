package tree

import (
	"fmt"
	"io"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  float64
}

type BinaryTree struct {
	root *TreeNode
}

func (t *BinaryTree) insert(data float64) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *TreeNode) insert(data float64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func multiplyNode(node *TreeNode) float64 {
	if node.left == nil && node.right == nil {
		return node.data
	} else {
		dl := multiplyNode(node.left)
		dr := multiplyNode(node.right)

		node.data = dl * dr
		return node.data
	}
}

func multiplyTree(tree *BinaryTree) {
	multiplyNode(tree.root)
}

func print(w io.Writer, node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}
