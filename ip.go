package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	instr := "a"

	fmt.Scan(&instr)
	if find := strings.Contains(instr, "."); find {
		// fmt.Println("find the character.")
		arr := strings.Split(instr, ".")
		if len(arr) == 4 {
			// fmt.Printf("ip arr %#v,len=%d\n", arr, len(arr))

			ip0, err := getipnum(arr[0])
			if err != nil {
				fmt.Println("not ip", arr)
				return
			}

			ip1, err := getipnum(arr[1])
			if err != nil {
				fmt.Println("not ip", arr)
				return
			}

			ip2, err := getipnum(arr[2])
			if err != nil {
				fmt.Println("not ip", arr)
				return
			}

			ip3, err := getipnum(arr[3])
			if err != nil {
				fmt.Println("not ip", arr)
				return
			}

			// fmt.Printf("ip = %d %d %d %d\n", ip0, ip1, ip2, ip3)
			intip := ip0<<24 + ip1<<16 + ip2<<8 + ip3
			fmt.Printf("%d\n", intip)
		} else {
			fmt.Printf("not ip arr %#v,len=%d\n", arr, len(arr))
		}

	} else {
		//fmt.Println("not find the character.")
		ipnum, err := strconv.Atoi(instr)
		if err != nil {
			fmt.Println("not ip", instr)
			return
		}
		ip0 := (ipnum >> 24)
		if ip0 > 255 {
			fmt.Printf("not ip %s,ip0=%d\n", instr, ip0)
			return
		}

		ip1 := (ipnum >> 16) & 0xff
		ip2 := (ipnum >> 8) & 0xff
		ip3 := ipnum & 0xff
		fmt.Printf("%d.%d.%d.%d\n", ip0, ip1, ip2, ip3)

	}
	//fmt.Printf("%s\n", instr)
}

func getipnum(ip string) (int, error) {
	ipnum, err := strconv.Atoi(ip)
	if err != nil || ipnum > 255 || ipnum < 0 {
		fmt.Println("not ip", ip)
		return 0, err
	}

	return ipnum, nil

}
