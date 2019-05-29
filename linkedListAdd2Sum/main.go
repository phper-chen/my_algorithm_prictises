package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 仅针对0-10之间的数字 1-9 头部不为0的链表
// 1->2->3->4
// 3->4->8->6->4
// 得到
// 4->6->1->1->5
type Lnode struct {
	V    int
	Next *Lnode
}

func (nodes *Lnode) createLnodes(lens int) {
	nodes.V = getRandomNum()
	cur := nodes
	for i := 1; i < lens; i++ {
		cur.Next = &Lnode{}
		cur.Next.V = getRandomNum()
		cur = cur.Next
	}
}
func (nodes *Lnode) printNodes() {
	var str string
	for temp := nodes; temp != nil; temp = temp.Next {
		str += strconv.Itoa(temp.V) + "->"
	}
	str = str[:len(str)-2]
	fmt.Println(str)
}
func getRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9) + 1
}
func addTwo(a, b *Lnode) *Lnode {
	if a == nil || a.Next == nil {
		return b
	}
	if b == nil || b.Next == nil {
		return a
	}
	l1, l2 := a, b
	newLnodes := &Lnode{} // 新的结果节点
	l := newLnodes
	sum := 0 // 每次两数相加之和
	c := 0   // sum超过9就得做进位 c用来记录进位
	for l1 != nil && l2 != nil {
		sum = l1.V + l2.V + c
		l.V = sum % 10
		c = sum / 10
		l1, l2 = l1.Next, l2.Next
		l.Next = &Lnode{}
		l = l.Next
	}
	if l1 == nil {
		for l2 != nil {
			sum = l2.V + c
			l.V = sum % 10
			c = sum / 10
			l2 = l2.Next
			if l2 != nil {
				l.Next = &Lnode{}
				l = l.Next
			}
		}
	}
	if l2 == nil {
		for l1 != nil {
			sum = l1.V + c
			l.V = sum % 10
			c = sum / 10
			l1 = l1.Next
			if l1 != nil {
				l.Next = &Lnode{}
				l = l.Next
			}
		}
	}
	if c == 1 {
		l.Next = &Lnode{V: c}
	}
	// 去掉头部的0
	for newLnodes.V == 0 {
		newLnodes = newLnodes.Next
	}
	return newLnodes
}
func main() {
	a := &Lnode{}
	a.createLnodes(12)
	a.printNodes()
	fmt.Println("第一个")
	b := &Lnode{}
	b.createLnodes(10)
	b.printNodes()
	fmt.Println("第二个")

	c := addTwo(a, b)
	c.printNodes()
}
