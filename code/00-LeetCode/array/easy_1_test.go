/*
 * @Author: duanhaobin
 * @Date: 2021-03-28 20:57:48
 */
package array

import (
	"fmt"
	"testing"
)

/*
	链接：https://leetcode-cn.com/problems/two-sum

	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

	你可以按任意顺序返回答案。

	示例 1：

	输入：nums = [2,7,11,15], target = 9
	输出：[0,1]
	解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

	解答：

	使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N)O(N) 降低到 O(1)O(1)。

	这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，
	即可保证不会让 x 和自己匹配。
*/

func TestTwoSum(t *testing.T) {

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

}
func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int, len(nums)-1)

	for i, v := range nums {
		if p, ok := hashTable[target-v]; ok {
			return []int{p, i}
		}
		hashTable[v] = i
	}

	return nil
}
