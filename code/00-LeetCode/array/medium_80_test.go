/*
 * @Author: duanhaobin
 * @Date: 2021-03-31 13:47:37
 */
package array

/*

	https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/

	给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。

	不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

*/
func removeDuplicates2(nums []int) int {
	fast, count := 1, 1

	for i, n := 1, len(nums); i < n; i++ {
		if nums[i] != nums[i-1] {
			count = 1
		} else {
			count++
		}
		if count <= 2 {
			nums[fast] = nums[i]
			fast++
		}
	}

	return fast
}
