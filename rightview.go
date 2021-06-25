package main

import "fmt"

type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var root *TreeNode = &TreeNode{Data: 6, Left: &TreeNode{"str", nil, nil}}
	res := RightView(root)
	fmt.Printf("main res %v\n", res)

}

func RightView(root *TreeNode) []interface{} {
	var res = []interface{}{}
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

	fmt.Printf("RightView res %v\n", res)
	return res

}
