package main

import "fmt"

type Lnode struct {
	V    int
	Next *Lnode
}

func createLnodes(node *Lnode, capacity int) {
	cur := node
	for i := 1; i < capacity; i++ {
		cur.Next = &Lnode{}
		cur.Next.V = i
		cur = cur.Next
	}
}

func printLnodes(node *Lnode) {
	for temp := node; temp != nil; temp = temp.Next {
		fmt.Println(temp)
	}
}
func insertReverse(node *Lnode) *Lnode {
	if node == nil || node.Next == nil {
		return nil
	}
	cur := node.Next
	// 解除头节点与后面的关系
	node.Next = nil
	var savePoint *Lnode
	for cur != nil {
		savePoint = cur.Next
		cur.Next = node
		// 该操作导致了node根节点的地址完全改变
		node = cur
		//		node.Next = cur.Next
		cur = savePoint
	}
	//	fmt.Println(node)
	//	fmt.Println(node.Next.V)
	//	fmt.Println(node.Next.Next.V)
	//	fmt.Println(node.Next.Next.Next.V)
	return node
}
func main() {
	node := &Lnode{}
	createLnodes(node, 9)
	printLnodes(node)
	a := insertReverse(node)
	printLnodes(a)
}
