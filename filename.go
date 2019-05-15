package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	fullFilename := "test.txt"
	fmt.Println("fullFilename =", fullFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename)
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	fmt.Println("fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenameOnly =", filenameOnly)
}
