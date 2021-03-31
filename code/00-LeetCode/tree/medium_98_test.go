/*
 * @Author: duanhaobin
 * @Date: 2021-03-30 21:49:13
 */
package tree

import "math"

/*

	分析：
	单纯的直接判断 leftNode.Val  < rootNode.Val && rightNode.Val > rootNode.Val 并不能判断成功

	因为
	- 左子树的所有节点的值都得小于root.Val
	- 右子树的所有节点的值都得大于root.Val

	因此，我们可以考虑使用上界、下界来帮我们判断。

	如果上界和下界都存在，判断当前节点的值是否在界内，如果不在界内返回false

	将当前节点的值作为所有左子树上界，继续对node.Left进行递归 ，这样就能判断所有左子树的节点值都能小于上界

	将当前节点的值作为所有右子树下界，继续对node.Right进行递归

	细节：
	- 空节点也是合理的二叉搜索树
*/

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
