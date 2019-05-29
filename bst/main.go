package main

import "fmt"

type BST struct {
	left  *BST
	value int
	right *BST
}

//查找元素
func (t *BST) Search(value int) bool {
	if t == nil {
		return false
	}
	compare := value - t.value
	if compare < 0 {
		return t.left.Search(value)
	} else if compare > 0 {
		return t.right.Search(value)
	} else {
		return true
	}
}

//添加元素
func (t *BST) Insert(value int) *BST {
	if t == nil {
		newNode := BST{nil, value, nil}
		return &newNode
	}
	if value < t.value {
		t.left = t.left.Insert(value)
	} else {
		t.right = t.right.Insert(value)
	}
	return t
}

/*删除元素
*1、如果被删除结点只有一个子结点，就直接将A的子结点连至A的父结点上，并将A删除
*2、如果被删除结点有两个子结点，将该结点右子数内的最小结点取代A。
 */
func (t *BST) Delete(value int) *BST {
	if t == nil {
		return t
	}
	compare := value - t.value
	if compare < 0 {
		t.left = t.left.Delete(value)
	} else if compare > 0 {
		t.right = t.right.Delete(value)
	} else { //找到结点,删除结点
		if t.left != nil && t.right != nil {
			t.value = t.right.getMin()
			t.right = t.right.Delete(t.value)
		} else if t.left != nil {
			t = t.left
		} else {
			t = t.right
		}
	}
	return t
}

//按顺序获得树中元素
func (t *BST) getAll() []int {
	values := []int{}
	return addValues(values, t)
}

//将一个节点加入切片中
func addValues(values []int, t *BST) []int {
	//	fmt.Printf("%v\n", values)
	if t != nil {
		values = addValues(values, t.left)
		values = append(values, t.value)
		values = addValues(values, t.right)
	}
	return values
}

//查找子树最小值
func (t *BST) getMin() int {
	if t == nil {
		return -1
	}
	if t.left == nil {
		return t.value
	} else {
		return t.left.getMin()
	}
}

//查找子树最大值
func (t *BST) getMax() int {
	if t == nil {
		return -1
	}
	if t.right == nil {
		return t.value
	} else {
		return t.right.getMax()
	}
}

//查找最小结点
func (t *BST) getMinNode() *BST {
	if t == nil {
		return nil
	} else {
		for t.left != nil {
			t = t.left
		}
	}
	return t
}

//查找最大结点
func (t *BST) getMaxNode() *BST {
	if t == nil {
		return nil
	} else {
		for t.right != nil {
			t = t.right
		}
	}
	return t
}

func main() {
	bsTree := BST{nil, 22, nil}
	bsTree.Insert(22)
	bsTree.Insert(12)
	bsTree.Insert(122)
	bsTree.Insert(62)
	bsTree.Insert(72)
	bsTree.Insert(25)
	bsTree.Insert(32)
	bsTree.Insert(42)
	fmt.Println(bsTree.getAll())

	fmt.Println(bsTree.Search(22))
	fmt.Println(bsTree.Search(20))

	fmt.Println(bsTree.getMin())
	fmt.Println(bsTree.getMinNode().value)

	fmt.Println(bsTree.getMax())
	fmt.Println(bsTree.getMaxNode().value)

	bsTree.Delete(22)
	fmt.Println(bsTree.getAll())

}
