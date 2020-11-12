package main

import (
	"fmt"
	"log"
)

type Pp struct {
	Aa int
	Bb int
}

func main() {
	ptest()
}

func ptest() {
	list := List{}
	list.Append(1)
	list.Append("a")
	list.ShowList()
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

func (l *List) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = l.headNode
	l.headNode = node
	return node
}

func (l *List) Append(data Object) {
	fmt.Printf("list append\n")
	node := &Node{Data: data}
	fmt.Printf("%#v\n", node)
	if l.IsEmpty() {
		l.headNode = node
	} else {
		cur := l.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
	l.ShowList()
}

func (l *List) Insert(index int, data Object) {
	if index < 0 {
		l.Add(data)
	} else if index > l.Length() {
		l.Append(data)
	} else {
		pre := l.headNode
		count := 0
		for count < (index - 1) {
			pre = pre.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

func (l *List) Remove(data Object) {
	pre := l.headNode
	if pre.Data == data {
		l.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

func (l *List) RemoveAtIndex(index int) {
	pre := l.headNode
	if index < 0 {
		l.headNode = pre.Next
	} else if index > l.Length() {
		log.Println("error,", index)
		return
	} else {
		count := 0
		for count != (index-1) && pre.Next != nil {
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

func (l *List) Contain(data Object) bool {
	cur := l.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		} else {
			cur = cur.Next
		}
	}
	return false
}

func (l *List) ShowList() {
	if !l.IsEmpty() {
		cur := l.headNode

		fmt.Printf("showlist")
		for {
			fmt.Printf("\t%#v", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("linkedlist empty\n")
	}
}
