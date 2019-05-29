package main

import (
	"fmt"
	"math/rand"
	"time"
)

//找到倒数第k个节点 使用快慢指针的思想来做 快指针比慢指针先行k步 然后双指针再一起遍历 当快指针遍历到的节点为nil时 慢指针所在位置便为倒数第k个节点

type Lnode struct {
	V    int
	Next *Lnode
}

func (nodes *Lnode) createNodes(lens int) {
	temp := nodes
	for i := 0; i < lens; i++ {
		temp.Next = &Lnode{V: getRandomNum()}
		temp = temp.Next
	}
}
func (nodes *Lnode) printNodes() {
	for temp := nodes.Next; temp != nil; temp = temp.Next {
		fmt.Println(temp.V)
	}
}
func getRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
func (nodes *Lnode) findKthNode(k int) *Lnode {
	// 初始化快慢指针
	slow, fast := nodes.Next, nodes.Next
	// 让快指针先行
	i := 0
	for ; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	// 说明查找的k已超出链表本身的长度
	if i < k {
		return nil
	}
	// 递进 如果快指针到头了 慢指针所在节点便是k的节点位置了
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main() {
	a := &Lnode{}
	a.createNodes(100)
	a.printNodes()
	k := 10
	c := a.findKthNode(k)
	if c != nil {
		fmt.Printf("倒数第%d个节点是%d", k, c.V)
	} else {
		fmt.Println("不存在该节点")
	}
}
