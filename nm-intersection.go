// 给定2个大小分别为n, m的整数集合，分别存放在两个数组中 int A[n], B[m]，输出两个集合的交集。
package main

import (
	"fmt"
)

func main() {
	var arrn = [3]int{1, 2, 3}
	var arrm = [4]int{1, 2, 3, 4}
	for _, vn := range arrn {
		for _, vm := range arrm {
			if vn == vm {
				fmt.Println(vn)
			}
		}
	}
}
