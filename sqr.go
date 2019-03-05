package main

import "fmt"

func main() {
	//fmt.Println("Hello, Worddld!")
	var sqr float64 = 0
	var input int64 = 78
	nn := 1
	for j:=0;j < 7;j++{
		fj0 := float64(nn)
		nn = 10*nn
		for i:=1;i < 11;i++{
			fi := float64(i)
			ii := fi/fj0
			sqrii := sqr + ii
			sqrii2 := sqrii * sqrii
			finput := float64(input)
			//fmt.Println(i,"fi=",fi,"j=",j,"fj0=",fj0,"ii=",ii,"sqrii=",sqrii,"sqrii2=",sqrii2,finput)
			if sqrii2 > finput {
				fmt.Println("i=",i,"j=",j)
				sqr = sqr + float64(i-1)/fj0
				fmt.Println(sqr)
				break
			}
		}
	}
	fmt.Println(input,"sqr",sqr)

}
 