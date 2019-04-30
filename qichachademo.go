package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type T struct {
	A int
	B string
}

func main() {

	fmt.Printf("this is main\n")

	key := ""
	seckey := ""
	t := strconv.FormatInt(time.Now().Unix(), 10)
	var str string = key + t + seckey
	fmt.Println("str=", str)
	fmt.Printf("dd=%s\n", t)

	h := md5.New()
	h.Write([]byte(str))                                                   // 需要加密的字符串为buf.String()
	fmt.Printf("tt=%s\n", strings.ToUpper(hex.EncodeToString(h.Sum(nil)))) // 输出加密结果

	c := &http.Client{
		Timeout: 300 * time.Second,
	}

	url := "https://api.qichacha.com/ECIV4/Search?key=" + key + "&dtype=json&keyword=abc"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Token", strings.ToUpper(hex.EncodeToString(h.Sum(nil))))
	reqest.Header.Add("Timespan", t)

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := c.Do(reqest)
	fmt.Println(response)

	f, err := os.Create("vlfile")
	if err != nil {
		fmt.Println("os Create", err.Error())
	}

	fmt.Println(f)

	io.Copy(f, response.Body)

	fmt.Println("http get ok")
	// token
	// curl  -H 'Token: ${tt}'  -H 'Timespan: ${dd}'  -X GET "https://api.qichacha.com/ECIV4/Search?key=34c13fc228904fac8352999e42a11ee0&dtype=json&keyword=abc"

}
