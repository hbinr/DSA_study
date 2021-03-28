/*
 * @Author: duanhaobin
 * @Date: 2021-03-28 19:42:24
 */
package dp

import (
	"fmt"
	"testing"
)

/*
	链接：https://leetcode-cn.com/problems/climbing-stairs

	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

	注意：给定 n 是一个正整数。

	示例 1：
	输入： 2
	输出： 2
	解释： 有两种方法可以爬到楼顶。
	1.  1 阶 + 1 阶
	2.  2 阶

	示例 2：
	输入： 3
	输出： 3
	解释： 有三种方法可以爬到楼顶。
	1.  1 阶 + 1 阶 + 1 阶
	2.  1 阶 + 2 阶
	3.  2 阶 + 1 阶

	n=1, 1种
	n=2, 2种
	n=3, 在n=1的基础上，爬2阶；在n=2的基础上，爬1阶，所以有 1+2 种
	n=4, 在n=2的基础上，爬2阶；在n=3的基础上，爬1阶， 有2+3

	以此类推，那么就是最后结果就是：f(n) = f(n-1)+f(n-2)

	这不就是斐波那契数列嘛？

	但是这道题，如果写一个纯递归的话，时间复杂度就是O(n^2)，非常大

	所以我们换种思路，尝试使用循环来做，自底向上来思考
*/

func TestClimbStairs(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("sum", sum)
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// 只保存最后的三个值，然后每次使用  f(n) = f(n-1)+f(n-2) 来计算最后一步
	// 同时更新阶梯树，自底向上
	f1, f2, f3 := 1, 2, 3

	// 此方法下索引是从1开始，不是0，所以 i<= n ，而不是 i < n
	for i := 3; i <= n; i++ {
		// 当 i = 3时，第一次循环，f3还是提前设置好的值，但还是算一遍f3 = 1+2 = 3，
		f3 = f1 + f2

		// 更新阶梯
		f1 = f2
		f2 = f3
	}

	return f3
}
