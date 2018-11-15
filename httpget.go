package main

import (
    "fmt"
    "net/http"
    "io"
    "os"
    "crypto/md5"
    "encoding/hex"
)


var (
    valuelinkurl = "http://47.92.255.245/valuelink.pdf"
    md5strsrc = "0b7d0d8c2b0109df9e5d0ba07b1028a3"
    vlfile = "vl.pdf"
)


func main() {
    //md5strsum := hex.EncodeToString(md5f.Sum(vlfile)
    checkMd5(valuelinkurl,vlfile,md5strsrc)

}

func checkMd5(url , infile , orgmd5 string){
    res,err := http.Get(url)
    if err != nil {
        fmt.Println("http get error",err.Error(),url)
    }

    f,err := os.Create(vlfile)
    if err != nil {
        fmt.Println("os Create",err.Error())
    }

    fmt.Println(f)

    io.Copy(f,res.Body)

    fmt.Println("http get ok")

    //md5
    md5f := md5.New()

    //f, _ := os.Open(vlfile)


    f.Seek(0,0)
    io.Copy(md5f,f)

    md5h := md5f.Sum(nil)
    md5strsum := hex.EncodeToString(md5h)

    fmt.Println("md5strsum",md5strsum,orgmd5)

    if orgmd5 == md5strsum {
        fmt.Println("md5 check OK",orgmd5)
    }

    fmt.Printf("%x\n",md5h)
    fmt.Printf("%x\n", md5f.Sum([]byte("")))

}
