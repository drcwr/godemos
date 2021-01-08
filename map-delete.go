package main

import (
	"log"
)

func initLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	initLog()

	maptest := make(map[string]chan struct{})

	maptest["del"] = make(chan struct{}, 1)

	log.Printf("maptest %#v\n", maptest)
	log.Printf("maptest 1 %#v\n", maptest["del"])
	cpu := maptest["dell"]
	log.Printf("maptest 1 2 %#v\n", cpu)
	// str := (&cpu).String()
	// log.Printf("maptest 1 3 %#v\n", str)

	ch, ok := maptest["del"]
	if ok {
		close(ch)
	} else {
		log.Printf("maptest 3 %#v\n", ch)
	}
	delete(maptest, "del")

	log.Printf("maptest 2 %#v\n", maptest["del"])
	ch, ok = maptest["del"]
	if ok {
		close(ch)
	} else {
		log.Printf("maptest 3 %#v\n", ch)
	}

}
