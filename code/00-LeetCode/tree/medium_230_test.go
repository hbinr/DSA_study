/*
 * @Author: duanhaobin
 * @Date: 2021-03-30 21:19:22
 */
package tree

/*
	https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

	给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）

	用中序遍历即可，并且每次都将k减1，当k为0时，当前节点的值即我们要求的值

	重点：二叉搜索树的中序遍历是有序的，搜索的结果是数组，第k小的元素对应数组的下标k-1
*/

func kthSmallest(root *TreeNode, k int) int {
	var res int
	var middleOrder func(node *TreeNode)

	middleOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		middleOrder(node.Left)

		// 每次都将k减1，当k为0时，当前节点的值即我们要求的值
		k -= 1
		if k == 0 {
			res = node.Val
			return
		}
		middleOrder(node.Right)
	}

	middleOrder(root)
	return res

}
