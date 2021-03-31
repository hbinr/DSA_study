/*
 * @Author: duanhaobin
 * @Date: 2021-03-31 13:47:22
 */
package array

/*
	https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
	给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

	不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。


*/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 0, 1
	for fast < n {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
