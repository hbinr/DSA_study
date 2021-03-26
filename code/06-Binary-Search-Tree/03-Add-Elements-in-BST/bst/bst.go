/*
 * @Author: duanhaobin
 * @Date: 2021-03-25 15:40:46
 */
package BST

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
	// tree 为空，设置根节点
	if b.root == nil {
		b.root = newNode(e)
		b.size++
	} else {
		b.add(b.root, e)
	}
}

// 向以node为根的二分搜索树中插入元素e，递归算法
func (b *BST) add(n *Node, e interface{}) {
	// 不处理重复数据的节点
	if n.e == e {
		return
		// 当前元素小于节点，则插入左子树
	} else if utils.Compare(e, n.e) < 0 && n.left == nil { // 左子树左子树递归终止条件
		n.left = newNode(e)
		b.size++
		return
		// 当前元素大于节点，则插入右子树
	} else if utils.Compare(e, n.e) > 0 && n.right == nil { // 右子树递归终止条件
		n.right = newNode(e)
		b.size++
		return
	}

	// 递归调用
	if utils.Compare(e, n.e) > 0 {
		b.add(n.left, e)
	} else {
		b.add(n.right, e)
	}
}
