package template

// 动态规划模板 : 一般适用于求最值

/*
	# 第一步：初识化base case
	dp[0][0] = base  case
	# 第二步：进行状态转移
	for 状态1 in 状态1的所有值:
		for 状态2 in 状态3的所有值:
			for ...
			# 第三步：套用表达式 (求最值)
			dp[状态1][状态2][...] = 求最值(选择1, 选择2, ...)

*/

// 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	// dp table
	dp := []int{}
	// 第一步：base case, 注意：dp存储了所有的状态，特殊案例中可能值需要dp table中的一部分
	dp[1], dp[2] = 1, 1

	// 第二步：状态转移
	for i := 3; i <= n; i++ {
		// 第三步：套用表达式
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
