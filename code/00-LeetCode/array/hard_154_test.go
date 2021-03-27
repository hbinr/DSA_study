/*
 * @Author: duanhaobin
 * @Date: 2021-03-27 18:01:58
 */
package array

/*
	2021-3-27  脉脉一面面试题 ，使用二分查找 O(log n)

	假设按照升序排序的数组在预先未知的某个点上进行了旋转。

	( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

	请找出其中最小的元素。

	注意数组中可能存在重复的元素。

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii


    不用mid和left比较 用mid和right比较的原因
	假如存在[1,2,3,4,5]和[2,3,4,5,1]这两种情况，左侧的mid和left比较，并不能决定最小值是在左侧还是右侧，因为左侧都是有序的。
	而如果比较mid和right的话，如果右侧是有序的，即nums[mid]<nums[right]，那最小值一定是在包含mid的左侧
*/
