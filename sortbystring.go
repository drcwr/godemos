package main

import (
	"log"
	"sort"
)

type T struct {
	name string
	tim  string
}

type TList []T

func main() {
	var ts = TList{{"n1", "3019-10-10T13:33:32+08:00"}, {"n2", "2019-10-10T14:33:32+08:00"}, {"n3", "4019-10-10T13:34:32+08:00"}, {"n4", "2019-11-10T13:33:32+08:00"}}
	t := ts[2]
	t.tim = t.tim[2:]
	ts = append(ts, t)
	log.Println("00000h,", ts)

	// Tsort(ts)
	// sort.Sort(ts)
	sort.Slice(ts, ts.Less)

	log.Println("switch,", ts)
}

func Tsort(ts TList) {
	for i := 0; i < len(ts); i++ {
		for j := i + 1; j < len(ts); j++ {
			if ts[i].tim < ts[j].tim {
				ts[i], ts[j] = ts[j], ts[i]
				log.Println(i, j)
			}
		}
		log.Println("ssssch,", ts)

	}

}

/*

type Interface interface {
    Len() int            // 获取元素数量
    Less(i, j int) bool // i，j是序列元素的指数。
    Swap(i, j int)        // 交换元素
}

*/

func (l TList) Len() int {
	return len(l)
}

func (l TList) Less(i, j int) bool {
	return l[j].tim < l[i].tim
}

func (l TList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// func Slice(slice interface{}, less func(i, j int) bool) {
