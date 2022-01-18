/*
 * @Author: duanhaobin
 * @Date: 2021-03-25 22:06:11
 */

package bst

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历：使用队列来实现，非递归
func LevelOrderDump(root *TreeNode) {
	if root == nil {
		return
	}
	levelStacks := []*TreeNode{root}

	for len(levelStacks) > 0 {
		// 出队
		branch := levelStacks[0]
		// 出队了元素要去掉
		levelStacks = levelStacks[1:]

		fmt.Print(branch.Val, ",")

		left := branch.Left
		right := branch.Right

		// 先左子树遍历
		if left != nil {
			levelStacks = append(levelStacks, left)
		}
		if right != nil {
			levelStacks = append(levelStacks, right)
		}
	}
	fmt.Println("")
}
