/*
 * @Author: duanhaobin
 * @Date: 2021-03-23 19:41:37
 * @LastEditTime: 2021-03-23 20:21:43
 */
package stack

import (
	"fmt"
	"testing"
)

func TestEasy20(t *testing.T) {

	fmt.Println("res: ", isValid("(){}"))
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 栈顶元素和给定括号相同，即'('匹配'('了，不满足，将该元素从栈顶删除
			// 这样设计，就保证栈顶始终是一个元素，节省了内存空间
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
