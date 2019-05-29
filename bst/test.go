package main

import (
	"fmt"
)

type BST struct {
	left  *BST
	v     int
	right *BST
}

func (t *BST) Search(v int) bool {
	if t != nil {
		return false
	}
	compare := v - t.v
	if compare < 0 {
		return t.left.Search(v)
	} else if compare > 0 {
		return t.right.Search(v)
	} else {
		return true
	}
}

func (t *BST) Insert(v int) *BST {
	if t == nil {
		newNode := BST{nil, v, nil}
		return &newNode
	}
	compare := v - t.v
	if compare < 0 {
		t.left = t.left.Insert(v)
	} else {
		t.right = t.right.Insert(v)
	}
	return t
}
func (t *BST) getMin() int {
	if t == nil {
		return nil
	}
	if t.left == nil {
		return t.v
	} else {
		return t.left.getMin()
	}
}

func (t *BST) getMax() int {
	if t == nil {
		return nil
	}
	if t.right == nil {
		return t.v
	} else {
		return t.right.getMax()
	}
}
func (t *BST) getMinNode() *BST {
	if t == nil {
		return nil
	}
	for t.left != nil {
		t = t.left
	}
	return t
}
func (t *BST) getMaxNode() *BST {
	if t == nil {
		return nil
	}
	for t.right != nil {
		t = t.right
	}
	return t
}

//1.如果删除的目标节点只有一个子树节点 则将子树提升 替代被删除的节点
//2.如果删除的目标节点有左右节点 那么就将右子树的最小值替换为被删除的节点值 然后再删除右子树的最小值节点
func (t *BST) Delete(v int) *BST {
	if t == nil {
		return nil
	}
	compare := t.v - v
	if compare < 0 {
		t = t.left.Delete(v)
	} else if compare > 0 {
		t.right = t.right.Delete(v)
	} else {
		if t.left != nil && t.right != nil {
			t.v = t.right.getMin()
			t.right = t.right.Delete(t.v)
		} else if t.left != nil {
			t = t.left
		} else {
			t = t.right
		}
	}
	return t
}

func (t *BST) GetAllValues() []int {
	values := []int{}
	return addValuesByMidSort(values, t)
}
func addValuesByMidSort(values []int, t *BST) []int {
	if t != nil {
		values = addValuesByMidSort(values, t.left)
		values = append(values, t.v)
		values = addValuesByMidSort(values, t.right)
	}
	return t
}
