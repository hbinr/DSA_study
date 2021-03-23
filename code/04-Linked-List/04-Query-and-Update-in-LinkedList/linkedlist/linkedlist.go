/*
 * @Author: duanhaobin
 * @Date: 2021-03-21 13:24:47
 * @LastEditTime: 2021-03-23 15:18:16
 */
package linkedlist

import (
	"bytes"
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
	dummyHead *Node // 虚拟头节点，不计入size
	size      int   // 链表大小
}

func New() *LinkedList {
	// 设置虚拟节点为空值
	return &LinkedList{
		dummyHead: &Node{},
	}

}

// 获取链表中的元素个数
func (l *LinkedList) GetSize() int {
	return l.size
}

// 返回链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("linkedlist : add failed, index is out of range")
	}
	// 头部直接使用现成的方法
	prev := l.dummyHead // 将头节点设置为prev，方便获取待插入节点的前一个节点

	// prev节点向前移动，获取待插入节点的前一个节点
	for i := 0; i < index; i++ { // 注意:不再从index为0的地方开始遍历，而是从虚拟头节点(index=0之前的节点)开始遍历
		prev = prev.next
	}

	// 插入新节点
	// node := &Node{
	// 	e:    e,
	// 	next: prev.next, // 1
	// }
	// prev.next = node // 2 注意：1，2顺序不能反了，否则新节点next便成了指向自己

	// 等同于上面5行代码
	prev.next = &Node{e, prev.next}
	l.size++
}

// 在链表头添加新元素e
func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

// 在链表末尾添加新的元素
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

// 获得链表的第index(0-based)个位置的元素
// 在链表中不是一个常用的操作，练习用
func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index > l.size {
		panic("linkedlist : Get failed, index is out of range")
	}

	cur := l.dummyHead.next // 虚拟节点不参与遍历，从使用下一个节点(真正头节点)开始
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

// 获得链表的第一个元素
func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

// 获得链表的最后一个元素
func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

// 修改链表的第index(0-based)个位置的元素为e
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("linkedlist : Get failed, index is out of range")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

// 查找链表是否存在元素e
func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next

	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}

	return false
}

// 输出链表内容
func (l *LinkedList) String() string {
	bf := bytes.Buffer{}
	cur := l.dummyHead.next

	for cur != nil {
		bf.WriteString(cur.String())
		bf.WriteString("->")
		cur = cur.next
	}

	// 最后再追加一个nil
	bf.WriteString("nil")

	return bf.String()
}
