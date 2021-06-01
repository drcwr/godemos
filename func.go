package main

import "fmt"

type funcT func(a ...interface{}) interface{}

func main() {
	funcagr(ff)
}

func funcagr(f funcT) {
	f("hehe")
}

func ff(input ...interface{}) interface{} {
	fmt.Println(input)
	return "test"
}
