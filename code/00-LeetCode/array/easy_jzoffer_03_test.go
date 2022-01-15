package array

import (
	"fmt"
	"testing"
)

/*
	https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
	找出数组中重复的数字。


	在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
	也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

	示例 1：

	输入：
	[2, 3, 1, 0, 2, 5, 3]
	输出：2 或 3
*/

func TestFindRepeatNumber(t *testing.T) {
	fmt.Println("res: ", findRepeatNumber2([]int{2, 3, 1, 0, 2, 5, 3}))

}

// 方法一：哈希表
func findRepeatNumber(nums []int) int {
	hashTable := make(map[int]bool)

	for _, num := range nums {
		if _, exists := hashTable[num]; exists {
			return num
		}
		hashTable[num] = true
	}
	return -1
}

// 方法二：原地置换
// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/mian-shi-ti-03-shu-zu-zhong-zhong-fu-de-shu-zi-yua/
// 思路：
// 1.可遍历数组并通过交换操作，使元素的 索引 与 值 一一对应（即 nums[i] = inums[i]=i ）。
// 因而，就能通过索引映射对应的值，起到与字典等价的作用。
// 2.关键点1：nums[i] == i,  如果相等，说明nums元素和索引一致，保证数组中 "索引=值" ，一一对应，那么i++，
// 3.关键点2：nums[nums[i]] == nums[i],  如果相等，说明有重复的元素，直接返回
// 4.关键点3：如果上述不相等，则需要置换元素位置，保证nums中的元素都是按顺序0,1,2,3......即保持和索引一致
func findRepeatNumber2(nums []int) int {
	i := 0

	for i < len(nums) {
		// 1. 保证数组元素=数组索引，一一对应
		if nums[i] == i {
			i++
			continue
		}
		// 2. 判重
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		// 3. 置换位置
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}

	return -1
}
