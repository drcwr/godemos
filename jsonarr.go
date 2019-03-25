package main

import (
    "fmt"
    "time"
    "encoding/json"
)

const (
    runs = 10000
)

type CompanyTrademark struct {
	RType int64 `json:"rType" form:"rType"`
	No string `json:"no" form:"no"`
	Image string `json:"image" form:"image"`
}

func main() {
    a := "啊啊啊"
    var ctm0 CompanyTrademark
    ctm0.RType = 1
    ctm0.Image = "imageimage"
    var ctm []CompanyTrademark
    var data []byte
    var err error

    fmt.Println("aaa len",len(a))
    start := time.Now()

    if data, err = json.Marshal(ctm0); err != nil {
        fmt.Println("CompanyTrademarkJson Marshal failed.")
        return
	}
    CompanyTrademarkStr := string(data)
    fmt.Println("CompanyTrademarkStr",CompanyTrademarkStr)
    str := fmt.Sprintf("[%s,%s]",CompanyTrademarkStr,CompanyTrademarkStr)

    fmt.Println("str",str)
    // str [{"rType":1,"no":"","image":"imageimage"},{"rType":1,"no":"","image":"imageimage"}]

    err = json.Unmarshal([]byte(str),&ctm)
    if err != nil {
        errMsg := "json.Unmarshal CompanyTrademarkJSON err"
        fmt.Println("errMsg",errMsg)
    }
    fmt.Println("ctm",ctm)



   duration := time.Now().Sub(start)
    fmt.Printf("Go: %.2fs\n", duration.Seconds())
}

