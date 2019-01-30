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
    // The Complete Works of William Shakespeare by William Shakespeare
    // http://www.gutenberg.org/files/100/100-0.txt
    fname := `100-0.txt` // "data.php"
    testdata := readFile(fname)
    needle := "Means to immure herself and not be seen."
    start := time.Now()

    for i := 0; i < runs; i++ {
        _ = strings.Contains(testdata, needle)

    }
    duration := time.Now().Sub(start)
    fmt.Printf("Go: %.2fs\n", duration.Seconds())

    fmt.Println(strings.Contains(testdata, needle))
    fmt.Println(strings.Index(testdata, needle))

}

func readFile(fname string) string {
    data, err := ioutil.ReadFile(fname)
    if err != nil {
        log.Fatal(err)
    }
    return string(data)
}
