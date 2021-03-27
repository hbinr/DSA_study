/*
 * @Author: duanhaobin
 * @Date: 2021-03-19 14:42:49
 * @LastEditTime: 2021-03-19 14:53:00
 * @FilePath: \DSA_study\code\00-LeetCode\tree\01-invert-binary-tree-test.go
 */
package tree

import (
	"testing"
)

/*
	翻转一棵二叉树。
	难度:easy

	示例：

	输入：

		4
	/   \
	2     7
	/ \   / \
	1   3 6   9
	输出：

		4
	/   \
	7     2
	/ \   / \
	9   6 3   1
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/invert-binary-tree
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestInvertBinaryTree(t *testing.T) {
	invertTree(nil)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Right = left
	root.Left = right
	return root
}
