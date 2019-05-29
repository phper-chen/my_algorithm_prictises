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
	for temp := node; temp != nil; temp = temp.Next {
		fmt.Println(temp.Data)
	}
}

func RecursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newNode := RecursiveReverseChild(node.Next)
	//	fmt.Println(newHead.Next)
	// 7->8->7->8.....
	node.Next.Next = node // 倒序
	//	fmt.Println(node.Next.Next)
	// 8->7->nil
	node.Next = nil // 清除前置指向

	return newNode
}

func main() {
	node := &LNode{}
	CreateList(node, 8)
	a := RecursiveReverseChild(node)
	PrintNode(a)
}
