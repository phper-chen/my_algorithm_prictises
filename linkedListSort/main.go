//void quicksort(Linklist head, Linklist end){
//    if(head == NULL || head == end)             //如果头指针为空或者链表为空，直接返回
//        return ;
//    int t;
//    Linklist p = head -> next;                  //用来遍历的指针
//    Linklist small = head;
//    while( p != end){
//        if( p -> data < head -> data){      //对于小于轴的元素放在左边
//            small = small -> next;
//            t = small -> data;
//            small -> data = p -> data;
//            p -> data = t;
//        }
//        p = p -> next;
//    }
//    t = head -> data;                           //遍历完后，对左轴元素与small指向的元素交换
//    head -> data = small -> data;
//    small -> data = t;
//    quicksort(head, small);                     //对左右进行递归
//    quicksort(small -> next, end);
//}

package main

// ListNode 是题目预定义的数据类型
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := split(head)
	return merge(sortList(left), sortList(right))
}

// 从中间位置，切分 list
func split(head *ListNode) (left, right *ListNode) {
	// head.Next != nil
	// 因为， sortList 已经帮忙检查过了

	// fast 的变化速度是 slow 的两倍
	// 当 fast 到达末尾的时候，slow 正好在 list 的中间
	slow, fast := head, head
	var slowPre *ListNode

	for fast != nil && fast.Next != nil {
		slowPre, slow = slow, slow.Next
		fast = fast.Next.Next
	}

	// 斩断 list
	slowPre.Next = nil
	// 使用 slowPre 是为了确保当 list 的长度为 2 时，会均分为两个长度为 1 的子 list

	left, right = head, slow
	return
}

// 把已经排序好了的两个 list left 和 right
// 进行合并
func merge(left, right *ListNode) *ListNode {
	// left != nil , right != nil
	// 因为， sortList 已经帮忙检查过了

	cur := &ListNode{}
	headPre := cur
	for left != nil && right != nil {
		// cur 总是接上较小的 node
		if left.Val < right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}

	// 处理 left 或 right 中，剩下的节点
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return headPre.Next
}
