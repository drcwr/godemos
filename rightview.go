package main

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root *TreeNode
	root = &TreeNode{Data: 6, Left: &TreeNode{Data: 7, Left: &TreeNode{Data: 9, Left: nil, Right: nil}, Right: nil}, Right: &TreeNode{Data: 8, Left: nil, Right: nil}}
	res := RightView(root)
	fmt.Printf("main res %v\n", res)
}

func RightView(root *TreeNode) []int {
	var res = []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Data)
	nextLayer := []*TreeNode{root}
	for len(nextLayer) > 0 {
		tmpLayer := []*TreeNode{}
		for k, v := range nextLayer {
			_ = k
			if v.Right != nil {
				tmpLayer = append(tmpLayer, v.Right)
			}
			if v.Left != nil {
				tmpLayer = append(tmpLayer, v.Left)
			}

		}
		if len(tmpLayer) > 0 {
			nextLayer = tmpLayer
			res = append(res, tmpLayer[0].Data)
		} else {
			break
		}

	}
	// res = append(res, []int{3, 4, 5}...)
	fmt.Printf("rightview res %v\n", res)
	return res
}
