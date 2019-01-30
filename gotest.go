package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "time"
)

const (
    runs = 10000
)

func main() {
    fname := "data.php"
    testdata := readFile(fname)
    needle := "576ad4f370014dfb1d0f17b0e6855f22"
    start := time.Now()

    for i := 0; i < runs; i++ {
        _ = strings.Contains(testdata, needle)

    }
    duration := time.Now().Sub(start)
    fmt.Printf("Go: %.2fs\n", duration.Seconds())
}

func readFile(fname string) string {
    data, err := ioutil.ReadFile(fname)
    if err != nil {
        log.Fatal(err)
    }
    return string(data)
}
