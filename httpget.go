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
    valuelinkurl = "https://avatars2.githubusercontent.com/u/5842099?s=400&u=674054525da4c81b7add2b68cf7544e73da9b936&v=4"
    md5strsrc = "c6958428d2c110ac2a863e1486b4844c"
    vlfile = "pig.jpg"
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
