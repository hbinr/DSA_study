package linkedlist

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-two-numbers
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 解决链接：https://leetcode-cn.com/problems/add-two-numbers/solution/jian-dan-yi-dong-javacpythonjs-pei-yang-y2w6g/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟头结点（构建新链表时的常用技巧，Val可以为任意值(满足数据类型)）
	dummy := &ListNode{Val: -1}
	// 记录进位
	cur, carry := dummy, 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x += l1.Val
		}

		if l2 != nil {
			y += l2.Val
		}

		// 加上进位
		total := x + y + carry
		// 计算下一个节点
		cur.Next = &ListNode{Val: total % 10}
		// cur 往前走，处理下一个节点
		cur = cur.Next

		// 处理进位
		carry = total / 10

		// 向后移动
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
