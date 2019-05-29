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

func RecursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	fmt.Println(newHead.Next)
	node.Next.Next = node // 倒序
	//	fmt.Println(node.Next.Next)

	node.Next = nil // 清除前置指向

	return newHead
}

func RecursiveReverse(node *LNode) {
	firstNode := node.Next
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}
func main() {
	node := &LNode{}
	CreateList(node, 8)
	RecursiveReverse(node)
	PrintNode(node)
}
