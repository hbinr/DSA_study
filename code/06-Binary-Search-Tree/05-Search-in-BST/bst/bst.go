/*
 * @Author: duanhaobin
 * @Date: 2021-03-25 16:49:09
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

// 看二分搜索树中是否包含元素 e
func (b *BST) Contains(e interface{}) bool {
	return b.contains(b.root, e)
}

// 看以 Node 为根的二分搜索树是否包含元素 e，递归算法
func (b *BST) contains(node *Node, e interface{}) bool {
	if node == nil {
		return false
	}
	if utils.Compare(e, node.e) == 0 {
		return true
	} else if utils.Compare(e, node.e) > 0 {
		return b.contains(node.right, e)
	} else {
		return b.contains(node.left, e)
	}
}
