// 给定字符串s, 要求把s中多于一个的连续空压缩成一个空格，并将连续的非空格字符串倒序打印出来，例如，给定"abc def efg"，打印"cba fed gfe"

package main

import "fmt"

func main() {
	deamon()
}

func deamon() {
	var s = "abc  def   efg         ddffggdd    kl"
	var tmp = []byte("")
	var total = []byte("")
	var lastchar = ' '
	for _, val := range s {
		// fmt.Printf("index=%d, val=%c\n", index, val)
		if val != ' ' {
			tmp = append(tmp, byte(val))
		} else {
			if lastchar != ' ' {
				// fmt.Printf("need daoxu str=%s---\n", tmp)
				daoxu(tmp)
				total = append(total, tmp...)
				tmp = tmp[0:0]
				total = append(total, byte(val))
				fmt.Printf("total str=%s---\n", total)
			}
		}
		lastchar = val
	}

	if len(tmp) > 0 {
		daoxu(tmp)
		total = append(total, tmp...)
	}
	fmt.Printf("end total str=%s---\n", total)

}

func daoxu(tmp []byte) {
	fmt.Printf("-- daoxu str=%s---\n", tmp)
	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
	}
	fmt.Printf("-- after daoxu str=%s---\n", tmp)

}
