package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

const (
	loop = 10000000 // 0
)

type CompanyTrademark struct {
	RType int64  `json:"rType" form:"rType"`
	No    string `json:"rNum" form:"rNum"`
	Image string `json:"rCer" form:"rCer"`
}

func Runstruct() int64 {
	var ctm0 CompanyTrademark
	ctm0.RType = 1
	ctm0.Image = "imageimage"
	start := time.Now()
	for i := 0; i < loop; i++ {
		json.Marshal(ctm0)
	}
	duration := time.Now().Sub(start)
	ret := duration.Nanoseconds()

	fmt.Printf("s Go: %d\n", ret)
	return ret
}

func Runmap() int64 {

	strstr := map[string]string{"rType": "0", "rNum": "nil", "rCer": "imageimage"}

	start2 := time.Now()
	for i := 0; i < loop; i++ {
		json.Marshal(strstr)
	}

	duration2 := time.Now().Sub(start2)
	ret := duration2.Nanoseconds()

	fmt.Printf("m Go: %d\n", ret)
	return ret

}

func runtest() {
	fmt.Println("runtest start")

	// time.Sleep(3 * time.Second)
	fmt.Println("runtest after sleep")
	retm := Runmap()
	_ = retm
	// rets := Runstruct()

	// fmt.Printf("m/s Go: %.6f\n", float64(retm)/float64(rets))
	fmt.Println("runtest exit")

}

func main() {

	fmt.Println("main start")

	go runtest()

	http.ListenAndServe("0.0.0.0:6060", nil)
	fmt.Println("main exit")

}
