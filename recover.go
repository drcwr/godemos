package main

import "log"

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	defer func() {
		err := recover()
		if err != nil {
			log.Println("recover err")
			return
		}
	}()

	log.Println("main")
}
