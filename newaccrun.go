package main

import (
    "fmt"
    _ "net/http/pprof"
    "os/exec"
    // "time"
    "io/ioutil"
)

func main() {

    // cmd := exec.Command("./apeiron" ,"new", "11111")
    cmd := exec.Command("./newacc")
    stdout, err := cmd.StdoutPipe()
    cmd.Start()
    content, err := ioutil.ReadAll(stdout)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(content))

}
