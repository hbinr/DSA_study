package array

/*
	LeetCode 两个题考察基本一模一样：
	https://leetcode-cn.com/problems/merge-sorted-array/

	https://leetcode-cn.com/problems/sorted-merge-lcci/submissions/
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 存放合并后的结果
	sorted := make([]int, 0, m+n)

	// p1:nums1的指针； p2:nums2的指针
	p1, p2 := 0, 0

	for {
		// 如果num1所有数据都先放到了sorted中，此时直接将 nums2 中剩余元素追加到sorted中
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}

		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		// 每次从两个数组头部取出比较小的数字放到结果中
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	// 将合并后的结果拷贝nums1中
	copy(nums1, sorted)
}
