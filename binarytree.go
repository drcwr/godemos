package main

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root *TreeNode
	root = &TreeNode{6, &TreeNode{Data: 7, Left: &TreeNode{Data: 9}}, &TreeNode{8, nil, nil}}
	fmt.Printf("\nPreOrder\n")
	PreOrder(root)
	fmt.Printf("\nMidOrder\n")
	MidOrder(root)
	fmt.Printf("\nAfterOrder\n")
	AfterOrder(root)
}

/*
PreOrder
MidOrder
AfterOrder

*/

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func MidOrder(root *TreeNode) {
	if root == nil {
		return
	}
	MidOrder(root.Left)
	fmt.Println(root.Data)
	MidOrder(root.Right)
}

func AfterOrder(root *TreeNode) {
	if root == nil {
		return
	}
	AfterOrder(root.Left)
	AfterOrder(root.Right)
	fmt.Println(root.Data)
}
