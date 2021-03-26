/*
 * @Author: duanhaobin
 * @Date: 2021-03-25 21:54:48
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

	if utils.Compare(e, node.e) > 0 {
		node.right = b.add(node.right, e)
	} else if utils.Compare(e, node.e) < 0 {
		node.left = b.add(node.left, e)
	}

	return node
}

// 前序遍历：使用栈来实现，非递归
func (b *BST) PreOrder() {
	stack := []*Node{}
	stack = append(stack, b.root)
	for len(stack) > 0 {
		// 出栈，获取当前节点的值
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(cur.e)

		// 先右子树遍历
		// 前序遍历不是先左子树遍历吗？此处是因为使用栈来遍历二叉树，后入先出
		if cur.right != nil {
			// 入栈
			stack = append(stack, cur.right)
		}

		// 后左子树遍历
		if cur.left != nil {
			stack = append(stack, cur.left)
		}
	}

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
