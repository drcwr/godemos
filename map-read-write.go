package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
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
	// for i := 0; i < 1000; i++ {
	// 	go read_map(maptestvar)

	// }
	go read_map(maptestvar)
	// go read_map(maptestvar)
	// go read_map(maptestvar)
	// go read_map(maptestvar)
	// go read_map(maptestvar)

	go write_map(maptestvar)

	for {
		time.Sleep(time.Second * 5)
	}

}

func read_map(m map[string]string) {
	for {
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		_, _ = m["testnil"]
		fmt.Println("read")
		// time.Sleep(time.Microsecond * 6)
	}
}

func write_map(m map[string]string) {
	for i := 0; ; i++ {
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		m["testnil"] = "www" + strconv.Itoa(i)
		fmt.Println("write")
		// time.Sleep(time.Microsecond * 6)
	}
}
