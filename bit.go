package main

import (
	"fmt"
	"log"
)

func main() {
	test()
}

func test() {
	var num int = 123123123123
	fmt.Printf("%b\n", num)

	t := num | (num >> 1)
	fmt.Printf("%b\n", t)

	t |= t >> 2
	fmt.Printf("%b\n", t)

	t |= t >> 4
	fmt.Printf("%b\n", t)

	t |= t >> 8
	fmt.Printf("%b\n", t)

	t |= t >> 16
	fmt.Printf("%b\n", t)

	t |= t >> 32
	fmt.Printf("%b,%d\n", t, t)
	countbit(num)
	countbit(t)
	countbit(0)
	countbit(1)
	countbit(2)
	countbit(3)
	countbit(4111)
}

func countbit(num int) int {
	var i = 0
	t := num
	for {
		if t == 0 {
			log.Println(i)
			break
		}
		i++
		t &= t - 1
	}
	return i
}
