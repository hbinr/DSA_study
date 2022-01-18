/*
 * @Author: duanhaobin
 * @Date: 2021-03-28 17:28:52
 */
package array

import (
	"fmt"
	"testing"
)

/*
     11. 盛最多水的容器

	使用双指针，从两个边界开始向中间收敛

	水容量：宽*高，即 (right - left)*min(height[left],height[right])

	1. 一开始，left =0, right = len(height)-1，此时宽度就是最宽的了
	2. 向中间移动，核心就是看柱子的高度，也就是height[i]的值，
	3. 当 height[left] < height[right]，应该left往中间移，即left++。
		为什么是left移动呢？因为目前宽度已经够宽了，主要看高了，如果height[left]小，那么需要看下一个柱子是否比height[left]大
		这样算出来的面积肯定当前的height[left]要大。
    4. 当 height[left] < height[right]，right往中间移，则right--。

	总之移动的原则就是，谁小移动谁，看下一个柱子是否够大，如果非常大的话，那么面积就会增加很多，毕竟宽度只是减少了1而已

	https://leetcode-cn.com/problems/container-with-most-water/submissions/
*/

func TestMaxArea(t *testing.T) {
	fmt.Println("maxArea: ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	max, left, right := 0, 0, len(height)-1

	for left < right {
		hLef, hRight := height[left], height[right]
		area := (right - left) * mini(hLef, hRight)
		if area > max {
			max = area
		}
		if hLef < hRight {
			left++
		} else {
			right--
		}
	}

	return max
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 更优雅的代码
func maxAreaPlus(height []int) int {
	area, max := 0, 0
	left, right := 0, len(height)-1

	for left < right {
		if hl, hr, dist := height[left], height[right], right-left; hl < hr {
			area = dist * hr
			left++
		} else {
			area = dist * hl
			right--
		}

		if area > max {
			max = area
		}
	}

	return max
}
