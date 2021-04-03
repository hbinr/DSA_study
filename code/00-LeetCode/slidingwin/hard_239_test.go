/*
 * @Author: duanhaobin
 * @Date: 2021-04-02 10:53:13
 */
package slidingwin

import (
	"fmt"
	"testing"
)

/*
	https://leetcode-cn.com/problems/sliding-window-maximum/

	239. 滑动窗口最大值

	给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。

	你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

	返回滑动窗口中的最大值。


*/

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println("res: ", maxSlidingWindow1([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

}

// 方法一：暴力解决 时间复杂度为：O(n^2)
func maxSlidingWindow1(nums []int, k int) []int {
	res := []int{}
	// 1. 先计算出第一个窗口的最大值
	tmpMax := max(nums[0:k])

	// 2. 遍历，为什么是 len - k + 1呢
	// 因为滑动窗口 i+k，所以计算到len - k时就相当于 遍历完整个数组了
	// 但是要注意 i从0开始，且i是小于len的，所以应该遍历 len(nums)-k+1，多了个+1
	for i := 0; i < len(nums)-k+1; i++ {
		// 3. 移动窗口i+k，并计算当前窗口中的最大值
		tmpMax = max(nums[i : i+k])
		// 4. 保存最大值
		res = append(res, tmpMax)
	}
	return res
}

func max(nums []int) int {
	max := nums[0]
	for i, n := 0, len(nums); i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 更优化的max写法
func maxInts(nums ...int) int {
	maxNum := -int(^uint(0)>>1) - 1
	fmt.Println("max Num", maxNum)
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
