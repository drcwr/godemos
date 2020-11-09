package main

import "log"

func main() {
	log.Println("max")

	var arr = []int{1, 2, 3, 4, 5, 6, 8, 10, 20, 30, 60, 51, 10, 5, 3}
	maxv := maxtatol(arr)
	log.Println("maxv=", maxv)
	maxr := max(maxtatol(arr[:len(arr)-1]), maxtatol(arr[1:]))
	log.Println("maxr=", maxr)
}

func maxtatol(arr []int) int {
	var v = [2]int{0, 0}
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	v = [2]int{arr[0], max(arr[0], arr[1])}
	for _, val := range arr[2:] {
		t := v[1]
		v[1] = max(v[0]+val, v[1])
		v[0] = t
		log.Println(v)
	}
	return max(v[0], v[1])
}

func max(n0, n1 int) int {
	if n0 > n1 {
		return n0
	}
	return n1

}
