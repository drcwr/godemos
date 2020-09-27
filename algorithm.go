package main

import (
	"log"
)

/*

algorithm.go demo

*/

func main() {
	var data = []int{1, 5, 7, 8, 9, 3, 10, 4, 6, 2, 11}

	initLog()

	log.Println("befor", data)
	// bobo(data)
	quicksort(data)
	log.Println("after", data)

	// gochan()

	return
}

func initLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func quicksort(data []int) {
	quick(data, 0, len(data)-1)
}

func quick(data []int, left, right int) {
	if left >= right {
		return
	}
	i, j, key := left, right, data[left]
	for i < j {
		for i < j && data[j] > key {
			j--
		}
		if i < j {
			data[i] = data[j]
			i++
		}
		for i < j && data[i] < key {
			i++
		}
		if i < j {
			data[j] = data[i]
			j--
		}
	}
	data[i] = key
	quick(data, left, i-1)
	quick(data, i+1, right)
}
