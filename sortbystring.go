package main

import (
	"log"
)

type T struct {
	name string
	tim  string
}

func main() {
	Tsort()
}

func Tsort() {
	var ts = []T{{"n1", "2019-10-10T13:33:32+08:00"}, {"n2", "2019-10-10T14:33:32+08:00"}, {"n3", "2019-10-10T13:34:32+08:00"}, {"n4", "2019-11-10T13:33:32+08:00"}}
	log.Println("00000h,", ts)
	for i := 0; i < len(ts); i++ {
		for j := i + 1; j < len(ts); j++ {
			if ts[i].tim < ts[j].tim {
				ts[i], ts[j] = ts[j], ts[i]
				log.Println(i, j)
			}
		}
		log.Println("ssssch,", ts)

	}

	log.Println("switch,", ts)
}
