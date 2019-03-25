package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "time"
    "regexp"
    "strconv"
)

const (
    runs = 10000
)

func main() {
    a := "啊啊啊"
    fmt.Println("aaa len",len(a))
    start := time.Now()

    isChinaIDCard("420881198601240031")

   duration := time.Now().Sub(start)
    fmt.Printf("Go: %.2fs\n", duration.Seconds())
}

func readFile(fname string) string {
    data, err := ioutil.ReadFile(fname)
    if err != nil {
        log.Fatal(err)
    }
    return string(data)
}

func isChinaIDCard(str string) bool {
	match, _ := regexp.Match(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`, []byte(str))

	if match {
		arr_int := []int{7,9,10,5,8,4,2,1,6,3,7,9,10,5,8,4,2}
		arr_ch := []string{"1","0","X","9","8","7","6","5","4","3","2"}
		sign := 0

		for k, v := range arr_int {
			int_temp, _ := strconv.Atoi(string([]byte(str)[k:k+1]))
			sign += int_temp * v
		}

		n := sign % 11
		val_num := arr_ch[n]
		if val_num != string([]byte(str)[17:18]) {
			return false
		}
	}
	return match
}
