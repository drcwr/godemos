package main

import (
	"fmt"
	"log"
	"reflect"
)

type Pp struct {
	Aa int
	Bb int
}

func main() {
	newtest()
}

func newtest() {
	var intp *int = new(int)
	var intsp = new([]int)
	log.Printf("newtest %p,%v,%p\n", intp, *intp, intsp)
}

func ptest() {
	var b int = 2
	var a *int = &b
	var c *int = nil
	var d = &c
	log.Println(c, d)

	t := reflect.TypeOf(*a)
	tt := reflect.TypeOf(t)
	log.Printf("%#v,%#v\n", t.Name(), tt.Kind())
}

func test() (*Pp, int, error) {
	var ret Pp = Pp{}
	// log.Printf("ret addr=%p\n", &ret)

	return &ret, 1, fmt.Errorf("test")
}

func test1() {
	i, k, err := test()
	log.Printf("i addr=%p,err add=%p,%#v\n", &i, &err, k)
	j, k, err := test()
	log.Printf("j addr=%p,err add=%p,%#v\n\n", &j, &err, k)
}

func test2() {
	i, k, err := test()
	log.Printf("i addr=%p,err add=%p,%#v\n", &i, &err, k)
	j, k, err1 := test()
	log.Printf("j addr=%p,err1 add=%p,%#v\n\n", &j, &err1, k)
}

func test3() {
	i, k, err := test()
	log.Printf("i addr=%p,err add=%p,%#v\n", &i, &err, k)
	i1, k, err := test()
	log.Printf("i addr=%p,err1 add=%p,%#v\n\n", &i1, &err, k)
}

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	} else {
		return false
	}
}

func (l *List) Length() int {
	cur := l.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func (l *List) add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = l.headNode
	l.headNode = node
	return node
}

func (l *List) Append(data Object) {
	node := &Node{Data: data}
	if l.IsEmpty() {
		l.headNode = node
	} else {
		cur := l.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}
