package main

import (
	"fmt"
)

type LNode struct {
	Data int
	Next *LNode
}

func CreateList(node *LNode, max int) {
	cur := node
	for i := 1; i <= max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}
func PrintNode(node *LNode) {
	for temp := node.Next; temp != nil; temp = temp.Next {
		fmt.Println(temp.Data)
	}
}
func InsertReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var next *LNode
	// head->1->2->3->4
	cur := node.Next.Next // 2
	node.Next.Next = nil  // 让1的指向为空 成为尾节点 不再指向2
	for cur != nil {
		// 把1放到2后面 同时记录2之后的指向 以便下次周期继续处理
		next = cur.Next
		// 把1放到2后面
		cur.Next = node.Next
		// 此时node为head->2->1
		node.Next = cur
		// 将后续未处理的的列表给cur继续处理
		cur = next
	}
}
func main() {
	node := &LNode{}
	CreateList(node, 8)
	InsertReverse(node)
	PrintNode(node)
}
