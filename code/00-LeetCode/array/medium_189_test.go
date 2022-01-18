package array

/*
	https://leetcode-cn.com/problems/rotate-array/

	给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。


	进阶：
	尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
	你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？


	示例 1:
	输入: nums = [1,2,3,4,5,6,7], k = 3
	输出: [5,6,7,1,2,3,4]

	解释:
	向右旋转 1 步: [7,1,2,3,4,5,6]
	向右旋转 2 步: [6,7,1,2,3,4,5]
	向右旋转 3 步: [5,6,7,1,2,3,4]

	解答：
	我们只需要将所有元素反转，然后反转前 k 个元素，再反转后面l-k个元素，就能得到想要的结果。

*/

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)     // 全部反转
	reverse(nums[:k]) // 反转前 k 个元素
	reverse(nums[k:]) // 再反转后面l-k个元素
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
