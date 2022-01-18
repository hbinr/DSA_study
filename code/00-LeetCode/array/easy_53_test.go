/*
 * @Author: duanhaobin
 * @Date: 2021-03-31 11:12:15
 */
package array

import (
	"fmt"
	"testing"
)

/*
	53.最大子序和

	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

思路
	1. 这道题用动态规划的思路并不难解决，比较难的是后文提出的用分治法求解，但由于其不是最优解法，所以先不列出来
	2. 动态规划的是首先对数组进行遍历，当前最大连续子序列和为 sum，结果为 ans
	3. 如果 sum > 0，则说明 sum 对结果有增益效果，则 sum 保留并加上当前遍历数字
	4. 如果 sum <= 0，则说明 sum 对结果无增益效果，需要舍弃，则 sum 直接更新为当前遍历数字
	5. 每次比较 sum 和 ans的大小，将最大值置为ans，遍历结束返回结果

	时间复杂度：O(n)O(n)

*/

func TestMaxSubArray(t *testing.T) {
	fmt.Println("res: ", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	sum, ans := 0, nums[0]

	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		ans = max2(sum, ans)
	}

	return ans
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
