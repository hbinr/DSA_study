package template

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表模板：迭代+递归

// 迭代模板
func traverseListNode(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		//	迭代遍历 p.Val
	}
}

// 递归模板
func traverseListNodeRe(head *ListNode) {

	// 前序遍历：在traverseRecursion(head.Next)调用之前操作head.Val
	traverseListNodeRe(head.Next)
	// 后序遍历：在traverseRecursion(head.Next)调用之前操作head.Val
}
