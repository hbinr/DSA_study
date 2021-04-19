package array

/*
	LeetCode 两个题考察基本一模一样：
	https://leetcode-cn.com/problems/merge-sorted-array/

	https://leetcode-cn.com/problems/sorted-merge-lcci/submissions/
*/
//  方法一
func merge1(nums1 []int, m int, nums2 []int, n int) {
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

//  方法二 从后到前把nums1的位置根据比较大小的方式填满，然后如果nums2还有部分剩余没有被填进去，那就逐个填进去
// https://leetcode-cn.com/problems/merge-sorted-array/solution/gu-du-ji-mo-jian-ji-golangshuang-100jie-0xpg0/
func merge2(nums1 []int, m int, nums2 []int, n int) {
	// i：nums1中 m个元素的最后一位
	// j：nums2中 n个元素的最后一位
	// cur：nums1最后一位
	i, j, cur := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		// 如果nums遍历完，并且nums[i] > nums[j]，那么将较大的数放到nums1后面(即cur索引位置)
		// 并且向前移动cur和i
		if nums1[i] > nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else { //
			nums1[cur] = nums2[j]
			j--
		}
		cur--
	}

	//当nums2未完全填充的时候
	for j >= 0 {
		nums1[cur] = nums2[j]
		cur--
		j--
	}
}
