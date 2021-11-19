package main

import (
	"archive/zip"
	"log"
)

func main() {
	fn := "test.zip"

	f, err := zip.OpenReader(fn)
	if err != nil {
		log.Printf("zip err:%v\n", err)
		panic(err.Error())
	}
	defer f.Close()
	log.Printf("zip f:%#v\n", f)
	// log.Printf("zip f:%#v\n", f.Reader.File[0].FileHeader.Name)
	// log.Printf("zip f:%#v\n", len(f.Reader.File))

	for i := 0; i < len(f.Reader.File); i++ {
		log.Printf("zip f.%d:%#v\n", i, f.Reader.File[i].FileHeader.Name)
	}
}
