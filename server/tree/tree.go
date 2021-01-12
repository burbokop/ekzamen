package tree

import (
	"fmt"
	"io"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Data  float64
}

type BinaryTree struct {
	Root *TreeNode
}

func (t *BinaryTree) Insert(data float64) *BinaryTree {
	if t.Root == nil {
		t.Root = &TreeNode{Data: data, Left: nil, Right: nil}
	} else {
		t.Root.Insert(data)
	}
	return t
}

func (n *TreeNode) Insert(data float64) {
	if n == nil {
		return
	} else if data <= n.Data {
		if n.Left == nil {
			n.Left = &TreeNode{Data: data, Left: nil, Right: nil}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Data: data, Left: nil, Right: nil}
		} else {
			n.Right.Insert(data)
		}
	}
}

func MultiplyNode(node *TreeNode) float64 {
	var dl float64
	var dr float64

	if node.Left != nil {
		dl = MultiplyNode(node.Left)
	} else {
		dl = 1
	}
	if node.Right != nil {
		dr = MultiplyNode(node.Right)
	} else {
		dr = 1
	}

	if node.Left == nil && node.Right == nil {
		return node.Data
	}
	node.Data = dl * dr
	return node.Data
}

func MultiplyTree(tree *BinaryTree) {
	MultiplyNode(tree.Root)
}

func Print(w io.Writer, node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Data)
	Print(w, node.Left, ns+2, 'L')
	Print(w, node.Right, ns+2, 'R')
}
