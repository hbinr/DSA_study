package array

import (
	"fmt"
	"testing"
)

/*
	https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/

	给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

	输入：nums = [1,3,5,4,7]
	输出：3
	解释：最长连续递增序列是 [1,3,5], 长度为3。
	尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。

	分析：
	使用记录当前连续递增序列的开始下标和结束下标，遍历数组的过程中每次比较相邻元素，
	根据相邻元素的大小关系决定是否需要更新连续递增序列的开始下标。

	具体而言，令 start 表示连续递增序列的开始下标，初始时 start=0，然后遍历数组nums，进行如下操作。
		1. 如果下标 i>0 且 nums[i]≤nums[i−1]，则说明当前元素小于或等于上一个元素

		因此 nums[i−1] 和 nums[i] 不可能属于同一个连续递增序列，必须从下标 i 处开始一个新的连续递增序列，

		因此令 start=i。如果下标 i=0 或 nums[i]>nums[i−1]，则不更新 start 的值。

    	2. 此时下标范围 [start,i] 的连续子序列是递增序列，其长度为 i−start+1，

		使用当前连续递增序列的长度更新最长连续递增序列的长度。




*/

func TestFindLengthOfLCIS(t *testing.T) {
	fmt.Println(findLengthOfLCIS([]int{1}))

}
func findLengthOfLCIS(nums []int) int {
	ans, start := 0, 0
	for i, n := 0, len(nums); i < n; i++ {
		// 必须保证 i > 0，这样数组才不会越界，出现-1情况
		if i > 0 && nums[i] <= nums[i-1] { // 说明当前元素小于或等于上一个元素
			// nums[i−1]和 nums[i] 不可能属于同一个连续递增序列，必须从下标i 处开始一个新的连续递增序列
			start = i
		}
		ans = max(ans, i-start+1) // 目标数组为：nums[start:i],那么该数组的长度为 i-start+1
	}
	fmt.Println("连续递增数组为: ", nums[start:ans])
	return ans
}
