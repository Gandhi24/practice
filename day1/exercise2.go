package main

import (
	"fmt"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func (root *Node) PreOrderTraversal() {
	if root != nil {
		fmt.Print(root.data)
		root.left.PreOrderTraversal()
		root.right.PreOrderTraversal()
	}
}

func (root *Node) PostOrderTraversal() {
	if root != nil {
		root.left.PostOrderTraversal()
		root.right.PostOrderTraversal()
		fmt.Print(root.data)
	}
}

func main() {
	root := Node{"+", nil, nil}
	root.left = &Node{"a", nil, nil}
	root.right = &Node{"-", nil, nil}
	root.right.left = &Node{"b", nil, nil}
	root.right.right = &Node{"c", nil, nil}

	fmt.Println("PreOrder traversal : ")
	root.PreOrderTraversal()
	fmt.Println("\n")
	fmt.Println("PostOrder traversal : ")
	root.PostOrderTraversal()
}
