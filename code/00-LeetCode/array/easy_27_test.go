package array

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	fmt.Println("res: ", removeElement([]int{3, 2, 2, 3}, 3))

}

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
