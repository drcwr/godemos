package main

// 0 1| 2 3|  4  5  6  7 2^n
// n R| l r| 2l 2r 3l 3r
import "fmt"

func main() {
	fmt.Printf("main binary heap\n")
}

func GetParent(root int) int {
	return root / 2
}

func GetLeft(root int) int {
	return 2 * root
}

func GetRight(root int) int {
	return 2*root + 1
}

func Sink() {

}

func Swim() {

}
