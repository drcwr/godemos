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
	No string `json:"rNum" form:"rNum"`
	Image string `json:"rCer" form:"rCer"`
}
// body {"status":1,"data":[{"name":"殷倩秋","user_id":"6170","mobile":"18916666807","QQ":"7300240"}],"info":"成功"}
// fmt.Println("CheckBusinessMobile name",resp.Data.([]interface{})[0].(map[string]interface{})["name"].(string))

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
    // str [{"rType":0,"rNum":null,"rCer":"imageimage"},{"rType":1,"rNum":"","rCer":"imageimage"}]

    err = json.Unmarshal([]byte(str),&ctm)
    if err != nil {
        errMsg := "json.Unmarshal CompanyTrademarkJSON err"
        fmt.Println("errMsg",errMsg)
    }
    fmt.Println("ctm",ctm)

    strstr := `[{"rType":0,"rNum":null,"rCer":"imageimage"},{"rType":1,"rNum":"888666","rCer":"imageimage"}]`

    err = json.Unmarshal([]byte(strstr),&ctm)
    if err != nil {
        errMsg := "2 json.Unmarshal CompanyTrademarkJSON err"
        fmt.Println("errMsg",errMsg)
    }
    fmt.Println("strstr ctm",ctm)


   duration := time.Now().Sub(start)
    fmt.Printf("Go: %.2fs\n", duration.Seconds())
}

