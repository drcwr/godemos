package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "6519452640458211328"
	m := IsSnowFlake(str)
	fmt.Printf("m=%#v", m)

}

func IsSnowFlake(str string) bool {
	match, _ := regexp.Match(`(\d{18}|\d{19})`, []byte(str))
	return match
}
