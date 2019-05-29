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

func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *LNode // 记录后续节点
	var pre *LNode // 已经逆序的
	next := node.Next
	for next != nil {
		cur = next.Next // 记录后续节点
		next.Next = pre // 逆序
		pre = next      // 已经逆序的记录下来
		next = cur      // 指向后续节点
	}
	node.Next = pre // 头部以后的节点改为逆序过的
}
func main() {
	node := &LNode{}
	CreateList(node, 8)
	Reverse(node)
	PrintNode(node)
}
