/*
 * @Author: duanhaobin
 * @Date: 2021-03-31 18:30:40
 */
package tree

/*
	https://leetcode-cn.com/problems/symmetric-tree/solution/dui-cheng-er-cha-shu-by-leetcode-solution/

	我们采用递归的方式

	根据子树是否为空分为三个判断：
	1. 两子树都为空，则对称
	2. 一个为空，一个非空，则不对称
	3. 两个都非空，那么需要比较
		- 这两个子树的根节点值必须相同
		- 每个树的右子树都与另一个树的左子树镜像对称

	时间复杂度：O(n)
	空间复杂度：O(n)
*/
// 方法一：递归
func isSymmetric(root *TreeNode) bool {
	return helper2(root, root)
}

func helper2(left, right *TreeNode) bool {
	// 两子树都为空 则对称
	if left == nil && right == nil {
		return true
	}
	// 一个空，一个非空，则不对称
	if left == nil || right == nil {
		return false
	}
	// 都非空：
	// 1. 这两个子树的根节点值必须相同
	// 2. 每个树的右子树都与另一个树的左子树镜像对称
	return left.Val == right.Val &&
		helper2(left.Left, right.Right) &&
		helper2(left.Right, right.Left)
}

// 方法二：迭代
// 初始化时我们把根节点入队两次。每次提取两个结点并比较它们的值（队列中每两个连续的结点应该是相等的，而且它们的子树互为镜像）
// 然后将两个结点的左右子结点按相反的顺序插入队列中。
// 当队列为空时，或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，该算法结束。

func isSymmetric2(root *TreeNode) bool {

	left, right := root, root
	queue := []*TreeNode{}
	queue = append(queue, left)
	queue = append(queue, right)

	for len(queue) > 0 {
		left, right = queue[0], queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left)
		queue = append(queue, right.Right)

		queue = append(queue, left.Right)
		queue = append(queue, right.Left)
	}
	return true
}
