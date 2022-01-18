package main

// 二分查找法，在有序数组arr中，查找target
// 如果找到target，返回相应的索引index
// 如果没有找到target，返回-1
func binarySearch(arr []int64, n, target int64) int64 {
	var l, r int64
	l = 0
	r = n - 1 // 在arr[l,r]之中查找target，注意：此处是左闭右闭的区间，因为r为 n-1

	// 不满足条件的直接返回-1
	if target < arr[l] || target > arr[r] || l > r {
		return -1
	}

	for l <= r {

		mid := (l + r) / 2
		if arr[mid] == target {
			return mid
		}

		// 如果目标小于分界点
		if target < arr[mid] {
			// 则在arr[l,mid-1]之中查找
			r = mid - 1
		} else { // 如果目标大于分界点
			// 则在arr[mid+1,r]之中查找
			l = mid + 1
		}
	}
	return -1
}

func main() {

}
