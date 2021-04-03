package template

// 二分查找模板
func binarySearch(nums []int, target int) int {
	res := 0
	//  1. 定义双指针
	left, right := 0, len(nums)

	for left < right { // 此为通用条件，具体循环条件视问题而定
		// 2. 设置mid索引
		mid := left + (right-left)>>1

		// 3. 中间值和目标比较逻辑
		if nums[mid] == target {
			// TODO
		} else if nums[mid] < target {
			left++
		} else if nums[mid] < target {
			right++
		}
	}
	return res
}
