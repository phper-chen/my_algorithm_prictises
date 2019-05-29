package main

import (
	"fmt"
)

type node struct {
	val         int
	left, right *node
}

func createBinaryTree(arr []int) []node {
	tree := make([]node, 0)
	// 先赋值 制造一个无节点关系的且以节点为元素的数组
	for k, v := range arr {
		tree = append(tree, node{})
		tree[k].val = v
	}
	// 每次遍历两个节点 分别作为左节点和右节点 所以长度应该对半分
	for i := 0; i < len(arr)/2; i++ {
		tree[i].left = &tree[i*2+1]
		if i*2+2 < len(tree) {
			tree[i].right = &tree[i*2+2]
		}
	}
	return tree
}
func (node *node) PreOrder() {
	if node == nil {
		return
	}
	fmt.Printf("%d  ", node.val)
	node.left.PreOrder()
	node.right.PreOrder()
}
func (node *node) MidOrder() {
	if node == nil {
		return
	}
	node.left.MidOrder()
	fmt.Printf("%d  ", node.val)
	node.right.MidOrder()
}
func (node *node) AfterOrder() {
	if node == nil {
		return
	}
	node.left.AfterOrder()
	node.right.AfterOrder()
	fmt.Printf("%d  ", node.val)

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tree := createBinaryTree(arr)
	fmt.Printf("%v\n", tree)
	tree[0].PreOrder()
	fmt.Println()
	tree[0].MidOrder()
	fmt.Println()
	tree[0].AfterOrder()
}

////层次遍历(广度优先遍历)
//func (node *Node) BreadthFirstSearch() {
//	if node == nil {
//		return
//	}
//	result := []int{}
//	nodes := []*Node{node}
//	for len(nodes) > 0 {
//		curNode := nodes[0]
//		nodes = nodes[1:]
//		result = append(result, curNode.Value)
//		if curNode.Left != nil {
//			nodes = append(nodes, curNode.Left)
//		}
//		if curNode.Right != nil {
//			nodes = append(nodes, curNode.Right)
//		}
//	}
//	for _, v := range result {
//		fmt.Print(v, " ")
//	}
//}

////层数(递归实现)
////对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1
//func (node *Node) Layers() int {
//	if node == nil {
//		return 0
//	}
//	leftLayers := node.Left.Layers()
//	rightLayers := node.Right.Layers()
//	if leftLayers > rightLayers {
//		return leftLayers + 1
//	} else {
//		return rightLayers + 1
//	}
//}
