package main

import "fmt"

type LNode struct {
	Data int
	Next *LNode
}

func createList(node *LNode, max int) {
	cur := node
	for i := 1; i <= max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

func main() {
	nodes := &LNode{}
	createList(nodes, 9)
	printNode(nodes)

	insertReverse(nodes)
	//	printNode(a)
}

func insertReverse(node *LNode) *LNode {
	remain := &LNode{Data: node.Data, Next: nil}
	if node == nil || node.Next == nil {
		return remain
	}
	cur := node.Next
	var savePoint *LNode
	var temp *LNode
	for cur != nil {
		savePoint = cur.Next
		temp = cur
		temp.Next = remain
		remain = temp
		cur = savePoint
	}
	return remain
}

func printNode(node *LNode) {
	for temp := node; temp != nil; temp = temp.Next {
		fmt.Println(temp.Data)
	}
}
