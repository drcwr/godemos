package main

import (
	"fmt"
	"strconv"
)

func main() {
	var t = make([][][4]int, 9)
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				for l := 1; l < 4; l++ {
					// fmt.Printf("%d.%d.%d.%d--%d\n", i, j, k, l, i+j+k+l)
					t[i+j+k+l-4] = append(t[i+j+k+l-4], [4]int{i, j, k, l})
				}
			}
		}
	}

	var input = "12561211"
	var bytea = []byte(input)
	var ipstr = [4][]byte{}
	var i = len(input)
labeltotal:
	for k, v := range t[i-4] {
		fmt.Printf("%d-%d,%d,%d,%d----", k, v[0], v[1], v[2], v[3])
		ipstr[0] = bytea[0:v[0]]
		ipstr[1] = bytea[v[0] : v[0]+v[1]]
		ipstr[2] = bytea[v[0]+v[1] : v[0]+v[1]+v[2]]
		ipstr[3] = bytea[v[0]+v[1]+v[2]:]
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

func checkipnum(input []byte) bool {
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
