/*
 * @Author: duanhaobin
 * @Date: 2021-03-27 21:20:20
 */
package array

import (
	"fmt"
	"testing"
)

/*
	假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。

	请找出其中最小的元素。

	示例 1：

	输入：nums = [3,4,5,1,2]
	输出：1
	示例 2：

	输入：nums = [4,5,6,7,0,1,2]
	输出：0
	示例 3：

	输入：nums = [1]
	输出：1


	提示：
	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
*/
func TestFindMin(t *testing.T) {
	fmt.Println("min value: ", findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1

	if nums[right] > nums[0] {
		return nums[0]
	}

	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else { //  nums[mid]  < nums[right]
			right = mid
		}
	}

	return nums[left]
}
