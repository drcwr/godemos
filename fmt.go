package main

import (
    "fmt"
    _ "net/http/pprof"
    "strings"
    // "os/exec"
    // "time"
    // "io/ioutil"
    // "github.com/apeiron/base"
    // "github.com/apeiron/shell/account"
)



func getMD5FromUrl(url string)string {
    // url = "http://ip/downloadfile?xxxx"
    strlist := strings.Split(url, "?")
    fmt.Println(strlist)
    if len(strlist) > 1 {
        return strlist[1]
    }

    return ""
    
}

func main() {
    fmt.Printf("aaa\n")
    fmt.Printf("bbb\n")
    fmt.Println([]byte("ABCXYZabcxyz"))


    _ = getMD5FromUrl("12?34?56")
    // fmt.Println()


    // account, _ := account.NewAccount(password)
    // address := base.BytesToHex(account.Address)
    // fmt.Printf("新建账户:%s\n", address)



}
