/*
 * @Author: duanhaobin
 * @Date: 2021-03-28 21:38:34
 */
package array

import (
	"fmt"
	"testing"
)

/*
	https://leetcode-cn.com/problems/trapping-rain-water/

	接雨水

	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

*/

func TestTrap(t *testing.T) {
	fmt.Println("res: ", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ans, left, right := 0, 0, len(height)-1

	// l_max代表height[0...left]中最高柱子的高度
	// r_max代表height[right...len(height)-1]中最高柱子的高度
	l_max, r_max := height[left], height[right]

	for left < right {
		l_max = max(l_max, height[left])
		r_max = max(r_max, height[right])

		// ans += min(l_max, r_max) - height[i]
		if l_max < r_max {
			ans += l_max - height[left]
			left++
		} else {
			ans += r_max - height[right]
			right--
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
