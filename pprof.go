package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    // _ "net/http/pprof"
    "runtime/pprof"
    "flag"
    "os"
    "os/signal"
    "syscall"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

    flag.Parse()

    if *cpuprofile != "" {
        f,err := os.Create(*cpuprofile)
        if err != nil {
            fmt.Println("openfile error")
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }

    go forloop()

    signalChan := make(chan os.Signal, 1)
    // signal.Notify(signalChan, os.Interrupt)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        for _ = range signalChan {
            fmt.Println("\n收到终端信号，停止服务... \n")
            pprof.StopCPUProfile()
            os.Exit(0)
        }
    }()

    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func forloop() {
    i := 0
    j := 0
    for {
        i ++
        fmt.Println("hello world",i,j)
        if i > 100 {
            time.Sleep(1*time.Second)
            i = 0
            j++
        }
        // if j > 20 {
        //     pprof.StopCPUProfile()
        //     os.Exit(0)
        // }
        
    }
}
