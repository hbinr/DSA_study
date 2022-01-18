/*
 * @Author: duanhaobin
 * @Date: 2021-03-25 17:01:33
 */
package bst

import (
	"bytes"
	"fmt"

	"dsa.study/code/utils"
)

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

// 向二分搜索树中添加新的元素 e
func (b *BST) Add(e interface{}) {
	b.root = b.add(b.root, e)
}

// 向以 Node 为跟的二分搜索树中插入元素 e，递归算法
// 返回插入新节点后二分搜索树的根
func (b *BST) add(node *Node, e interface{}) *Node {
	if node == nil {
		b.size++
		return newNode(e)
	}

	// 当前节点的元素小于新元素 e，则向右子树中添加
	if utils.Compare(e, node.e) > 0 {
		node.right = b.add(node.right, e)
	} else if utils.Compare(e, node.e) < 0 {
		node.left = b.add(node.left, e)
	}

	return node
}

// 中序遍历
func (b *BST) InOrder() {
	b.inOrder(b.root)
}

// 中序遍历
func (b *BST) inOrder(node *Node) {
	if node == nil {
		return
	}
	b.inOrder(node.left)
	fmt.Println(node.e) // 中序遍历：在左右子树中间访问节点
	b.inOrder(node.right)
}

// 后序遍历
func (b *BST) PostOrder() {
	b.postOrder(b.root)
}

// 后序遍历
func (b *BST) postOrder(node *Node) {
	if node == nil {
		return
	}
	b.postOrder(node.left)
	b.postOrder(node.right)
	fmt.Println(node.e) // 后序遍历：在左右子树最后访问节点

}

// 打印二叉树
func (b *BST) String() string {
	var buffer bytes.Buffer
	generateBSTSting(b.root, 0, &buffer)
	return buffer.String()
}

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(n *Node, depth int, buffer *bytes.Buffer) {
	if n == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(n.e) + "\n")
	generateBSTSting(n.left, depth+1, buffer)
	generateBSTSting(n.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
