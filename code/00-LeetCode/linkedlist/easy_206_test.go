/*
 * @Author: duanhaobin
 * @Date: 2021-03-27 23:17:01
 */
package linkedlist

import "testing"

/*

	反转一个单链表。

	示例:

	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL

	进阶:
	你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

	链接：https://leetcode-cn.com/problems/reverse-linked-list

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestReverseLis(t *testing.T) {

}

// 方法一：迭代
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		// 1.保存cur的下一个节点
		tmp := cur.Next
		// 2.prev和cur交换位置，也就是说cur的next指针执行了prev
		cur.Next = prev
		// 3.向前移动prev
		prev = cur
		// 4.和cur
		cur = tmp

	}
	// 5.头节点为 prev
	return prev
}

// 方法一：递归
func reverseList2(head *ListNode) *ListNode {
	// 如果链表只有一个节点，反转也是它自己，直接返回
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
