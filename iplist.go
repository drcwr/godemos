package main

import "fmt"

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
	// for i := 0; i < len(t); i++ {
	// fmt.Printf("%d,len=%d\nt:%v\n\n", i+4, len(t[i]), t[i])
	// }

	var input = "123212122"
	var bytea = []byte(input)
	var i = len(input)
	for k, v := range t[i-4] {
		fmt.Printf("%d-%d,%d,%d,%d----", k, v[0], v[1], v[2], v[3])

		fmt.Printf("%v.", string(bytea[0:v[0]]))
		fmt.Printf("%v.", string(bytea[v[0]:v[0]+v[1]]))
		fmt.Printf("%v.", string(bytea[v[0]+v[1]:v[0]+v[1]+v[2]]))
		fmt.Printf("%s", string(bytea[v[0]+v[1]+v[2]:]))
		fmt.Printf("\n")

	}

}
