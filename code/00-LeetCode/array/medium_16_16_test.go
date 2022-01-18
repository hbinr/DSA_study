package array

/*
	https://leetcode-cn.com/problems/sub-sort-lcci/
	给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。

	示例：

	输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
	输出： [3,9]

	从左往右扫描，寻找逆序对 (正序: 逐渐变大)，找到最后一对逆序对所处位置
		1. 记录扫描过的最大值
		2. 如果发现当前值大于最大值，则记录当前值的位置

	从右往左扫描，寻找逆序对 (正序: 逐渐变小)，找到最后一对逆序对所处位置
		1. 记录扫描过的最小值
		2. 如果发现当前值小于最小值，则记录当前值的位置
*/

func subSort(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	// 从左往右扫描，寻找逆序对 (正序: 逐渐变大)，找到最后一对逆序对所处位置
	max := nums[0]
	right := -1
	for i := 1; i < n; i++ {
		v := nums[i]
		if v >= max {
			max = v
		} else {
			right = i
		}
	}

	// 如果right 为 -1，那么该数组已经是有序的，直接返回
	if right == -1 {
		return []int{-1, -1}
	}

	// 从右往左扫描，寻找逆序对 (正序: 逐渐变小)，找到最后一对逆序对所处位置
	min := nums[n-1]
	left := -1
	for i := n - 2; i >= 0; i-- {
		v := nums[i]
		if v <= min {
			min = v
		} else {
			left = i
		}
	}
	return []int{left, right}
}
