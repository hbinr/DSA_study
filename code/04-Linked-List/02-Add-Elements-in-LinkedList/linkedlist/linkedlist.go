/*
 * @Author: duanhaobin
 * @Date: 2021-03-21 11:02:00
 * @LastEditTime: 2021-03-21 11:31:49
 */

package linkedlist

import (
	"fmt"
)

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

// LinkedList 链表结构体
type LinkedList struct {
	head *Node // 头节点
	size int   // 链表大小
}

// 获取链表中的元素个数
func (l *LinkedList) GetSize() int {
	return l.size
}

// 返回链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表头添加新元素e
func (l *LinkedList) AddFirst(e interface{}) {
	// node := &Node{
	// 	e:    e,
	// }
	// l.head = node

	// 以上4行代码等同于下面4行代码
	l.head = &Node{
		e:    e,
		next: l.head,
	}
	// 记得链表大小加1
	l.size++
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("linkedlist : add failed, index is out of range")
	}
	// 头部直接使用现成的方法
	if index == 0 {
		l.AddFirst(e)
	} else {
		prev := l.head // 将头节点设置为prev，方便获取待插入节点的前一个节点

		// prev节点向前移动，获取待插入节点的前一个节点
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// 插入新节点
		node := &Node{
			e:    e,
			next: prev.next, // 1
		}
		prev.next = node // 2 注意：1，2顺序不能反了，否则新节点next便成了指向自己

		l.size++
	}
}

// 在链表末尾添加新的元素
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}
