package main

import (
	"fmt"
)

func main() {
	var y int32 = 2

	for y <= 8 {
		fmt.Println(y)
		y = y + y
	}

}
