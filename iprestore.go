package main

import (
	"fmt"
	"strconv"
)

func maket() [][][4]int {
	t := make([][][4]int, 12-4+1)
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				for l := 1; l < 4; l++ {
					t[i+j+k+l-4] = append(t[i+j+k+l-4], [4]int{i, j, k, l})
				}
			}
		}
	}
	return t
}

func main() {
	var t = maket()

	var input = "125610211"
	var ipstr = [4]string{}
	var i = len(input)
labeltotal:
	for k, v := range t[i-4] {
		fmt.Printf("%d-%d,%d,%d,%d----", k, v[0], v[1], v[2], v[3])
		ipstr[0] = input[0:v[0]]
		ipstr[1] = input[v[0] : v[0]+v[1]]
		ipstr[2] = input[v[0]+v[1] : v[0]+v[1]+v[2]]
		ipstr[3] = input[v[0]+v[1]+v[2]:]
		for i := 0; i < 4; i++ {
			if !checkipnum(ipstr[i]) {
				continue labeltotal
			}

		}
		fmt.Printf("%v", string(ipstr[0]))
		for i := 1; i < 4; i++ {
			fmt.Printf(".%v", string(ipstr[i]))
		}

		fmt.Printf("\n")

	}

}

func checkipnum(input string) bool {
	if len(input) > 1 && input[0] == '0' {
		fmt.Printf("CHECK FALSE,0 start,%s\n", string(input))
		return false
	}
	numip, err := strconv.Atoi(string(input))
	if err != nil {
		fmt.Printf("CHECK FALSE,%s\n", err.Error())
		return false
	}
	if numip > 255 || numip < 0 {
		fmt.Printf("CHECK FALSE,numip %v\n", numip)
		return false
	}

	return true
}
