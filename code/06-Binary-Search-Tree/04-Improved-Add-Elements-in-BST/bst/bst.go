/*
 * @Author: duanhaobin
 * @Date: 2021-03-25 16:17:46
 */
package bst

import "dsa.study/code/utils"

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

func newNode(e interface{}) *Node {
	return &Node{e: e}
}

type BST struct {
	root *Node
	size int
}

func New() *BST {
	return &BST{}
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// 向二分搜索树中添加新的元素e
func (b *BST) Add(e interface{}) {
	b.root = b.add(b.root, e)
}

// 向以node为根的二分搜索树中插入元素e，递归算法
//
// 较之前的算法：去掉了判断节点是否为空的情况，毕竟节点空很正常
func (b *BST) add(node *Node, e interface{}) *Node {
	// 如果节点空，则新建节点
	if node == nil {
		b.size++
		return newNode(e)
	}

	// 递归调用
	if utils.Compare(e, node.e) > 0 { // 如果新增数据大于节点数据，则插入右子树中
		node.right = b.add(node.right, e)
	} else if utils.Compare(e, node) < 0 {
		node.left = b.add(node.left, e)
	}

	return node
}
