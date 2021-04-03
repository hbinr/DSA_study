package template

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树遍历模板：迭代+递归

// 迭代：使用栈结构
func traverseBST(root *TreeNode) {

}

// 递归
func traverseBSTRe(root *TreeNode) {

	// 前序遍历：在左右子树前面访问节点
	traverseBSTRe(root.Left)
	// 中序遍历：在左右子树中间访问节点
	traverseBSTRe(root.Right)
	// 后序遍历：在左右子树最后访问节点
}
