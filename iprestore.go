package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list()
	dfsmain()
}

func dfs(s string, begin int, path []string, res *[]string) {

	if len(path) > 4 {
		return
	}

	if begin == len(s) && len(path) == 4 {
		*res = append(*res, strings.Join(path, "."))
		return
	}

	for i := begin; i < len(s) && i < begin+3; i++ {
		sub := s[begin : i+1]
		if ok := checkipnum(sub); ok {
			path = append(path, sub)
			dfs(s, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func dfsmain() {
	var (
		input = "125610211"
		res   []string
		path  = []string{}
	)
	dfs(input, 0, path, &res)

	fmt.Printf("dfs  res %v\n", res)
}

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

func list() {
	var (
		t     = maket()
		input = "125610211"
		ipstr = [4]string{}
		i     = len(input)
		res   = []string{}
	)
labeltotal:
	for _, v := range t[i-4] {
		// fmt.Printf("%d-%d,%d,%d,%d----", k, v[0], v[1], v[2], v[3])
		ipstr[0] = input[0:v[0]]
		ipstr[1] = input[v[0] : v[0]+v[1]]
		ipstr[2] = input[v[0]+v[1] : v[0]+v[1]+v[2]]
		ipstr[3] = input[v[0]+v[1]+v[2]:]
		for i := 0; i < 4; i++ {
			if !checkipnum(ipstr[i]) {
				continue labeltotal
			}

		}

		// fmt.Printf("%s\n", strings.Join(ipstr[:], "."))

		res = append(res, strings.Join(ipstr[:], "."))

	}
	fmt.Printf("list res %v\n", res)

}

func checkipnum(input string) bool {
	if len(input) > 1 && input[0] == '0' {
		// fmt.Printf("CHECK FALSE,0 start,%s\n", string(input))
		return false
	}
	numip, err := strconv.Atoi(string(input))
	if err != nil {
		// fmt.Printf("CHECK FALSE,%s\n", err.Error())
		return false
	}
	if numip > 255 || numip < 0 {
		// fmt.Printf("CHECK FALSE,numip %v\n", numip)
		return false
	}

	return true
}
