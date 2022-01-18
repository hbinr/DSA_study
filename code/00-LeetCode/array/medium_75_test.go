package array

import (
	"testing"
)

/*
	https://leetcode-cn.com/problems/sort-colors/

	75. 颜色分类  又称荷兰国旗

	给定乱序数组,元素只有0,1,2这三种数字，求排序后的数组

	要求：时间复杂度为O(n), 空间复杂度为O(1)

	分析：
		三个指针：
			1. left 指向数组头部
			2. right  指向数组尾部
			3. cur  指向当前元素所处下标

		遍历整个数组：
			1. 遇到1，跳过，cur++
			2. 遇到0，和 left指针交换值，left++，cur++
			3. 遇到2，和 right 指针交换值，right--，
			4. 再循环一次，对cur 指针进行上述判断，直到 cur > right 为止
*/

func TestSortColors(t *testing.T) {
	sortColors([]int{2, 1, 0, 1, 0, 2})
}

func sortColors(nums []int) {
	cur, left, right := 0, 0, len(nums)-1
	for cur <= right {
		v := nums[cur]
		if v == 0 {
			nums[left], nums[cur] = nums[cur], nums[left]
			left++
			cur++
		} else if v == 1 {
			cur++
		} else {
			nums[right], nums[cur] = nums[cur], nums[right]
			right--
		}
	}
}
