/*
 * @Author: duanhaobin
 * @Date: 2021-03-29 17:10:33
 */

package array

import (
	"fmt"
	"testing"
)

/*
     https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/

	一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。
	要求时间复杂度是O(n)，空间复杂度是O(1)。

	[1, 1, 3, 3, 5, 6, 6, 7, 7]

*/
func TestSingleNumbers(t *testing.T) {
	fmt.Println("res: ", singleNumbers1([]int{1, 1, 3, 3, 5, 6, 6, 7, 7}))

}

// 方法一：使用hashTable 做临时缓存
func singleNumbers1(nums []int) []int {
	hashTable := map[int]int{}
	n := len(nums)
	res := []int{}
	for i := 0; i < n; i++ {
		if _, ok := hashTable[nums[i]]; ok {
			hashTable[nums[i]] += 1
		} else {
			hashTable[nums[i]] = 1
		}

	}

	for i := 0; i < n; i++ {
		v, _ := hashTable[nums[i]]
		if v == 1 {
			res = append(res, nums[i])

		}
	}
	return res
}

// 方法二：分组异或
func singleNumbers2(nums []int) []int {
	var a int
	for i := range nums {
		a ^= nums[i]
	}
	mask := a & (-a)
	res := make([]int, 2)
	for _, v := range nums {
		if (v & mask) == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
