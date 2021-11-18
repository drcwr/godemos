package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	runs = 10000
)

type CompanyTrademark struct {
	RType int64  `json:"rType" form:"rType"`
	No    string `json:"no" form:"no"`
	Image string `json:"image" form:"image"`
}

type CompanyTrademark2 struct {
	CompanyTrademark
	RType int64  `json:"rType" form:"rType"`
	No    string `json:"no" form:"no"`
	Image string `json:"image" form:"image"`
}

func main() {
	a := "啊啊啊"
	var ctm0 CompanyTrademark2
	ctm0.CompanyTrademark.RType = 10
	ctm0.CompanyTrademark.No = "haha0"
	ctm0.CompanyTrademark.Image = "imageimage0"
	ctm0.RType = 1
	// ctm0.No = "haha1"
	ctm0.Image = "imageimage1"
	var ctm CompanyTrademark2
	var data []byte
	var err error

	fmt.Printf("ctm0 %v", ctm0)
	fmt.Printf("ctm0.CompanyTrademark %v", ctm0.CompanyTrademark)

	fmt.Println("aaa len", len(a))
	start := time.Now()

	if data, err = json.Marshal(ctm0); err != nil {
		fmt.Println("CompanyTrademarkJson Marshal failed.")
		return
	}
	CompanyTrademarkStr := string(data)
	fmt.Println("CompanyTrademarkStr", CompanyTrademarkStr)

	err = json.Unmarshal([]byte(CompanyTrademarkStr), &ctm)
	if err != nil {
		errMsg := "json.Unmarshal CompanyTrademarkJSON err"
		fmt.Println("errMsg", errMsg)
	}
	fmt.Println("ctm", ctm)
	fmt.Printf("ctm %v", ctm)
	fmt.Printf("ctm.CompanyTrademark %v", ctm.CompanyTrademark)

	err = json.Unmarshal([]byte("{\"kk\":\"vv\"}"), &ctm)
	if err != nil {
		errMsg := "json.Unmarshal CompanyTrademarkJSON err"
		fmt.Println("errMsg", errMsg, err.Error())
	}

	duration := time.Now().Sub(start)
	fmt.Printf("Go: %.2fs\n", duration.Seconds())
}
