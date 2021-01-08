package main

import (
	"log"
)

func initLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type mapstr map[string]string

func main() {
	initLog()
	var maptestvar map[string]string
	log.Printf("maptestvar %#v\n", maptestvar)
	if maptestvar == nil {
		// panic: assignment to entry in nil map

		// maptestvar = make(map[string]string)
		maptestvar = mapstr{}
	}

	maptestvar["testnil"] = "nilvalue"

	maptest := make(map[string]chan struct{})
	log.Printf("maptest make %#v\n", maptest)

	maptest["del"] = make(chan struct{}, 1)

	log.Printf("maptest %#v\n", maptest)
	log.Printf("maptest 1 %#v\n", maptest["del"])

}
