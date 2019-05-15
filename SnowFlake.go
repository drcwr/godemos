package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "6519452640458211328"
	m := IsSnowFlake(str)
	fmt.Printf("str=%s,m=%#v\n", str, m)

	str = "65194526404582113"
	m = IsSnowFlake(str)
	fmt.Printf("str=%s,m=%#v\n", str, m)

	str = "65194526404582113a"
	m = IsSnowFlake(str)
	fmt.Printf("str=%s,m=%#v\n", str, m)

	str = "651945264045821130"
	m = IsSnowFlake(str)
	fmt.Printf("str=%s,m=%#v\n", str, m)

	str = "65194526404582113281"
	m = IsSnowFlake(str)
	fmt.Printf("str=%s,m=%#v\n", str, m)

}

func IsSnowFlake(str string) bool {
	match, _ := regexp.Match(`(\d{18}|\d{19})`, []byte(str))
	return match
}
