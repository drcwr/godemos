package main

import (
    "fmt"
)

func main() {
    var a [3]int
    fmt.Println(a)
    a[0] = 12
    a[1] = 22
    a[2] = 32
    fmt.Println(a)
    for i,v := range a {
        fmt.Println(i,v)
    }
}

