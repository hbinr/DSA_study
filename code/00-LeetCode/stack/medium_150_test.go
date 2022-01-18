/*
 * @Author: duanhaobin
 * @Date: 2021-03-24 10:57:09
 */
package stack

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMedium150(t *testing.T) {
	fmt.Println("res: ", evalRPN2([]string{"4", "13", "5", "/", "+"}))

}

// 方法一：栈顶保存最后计算结果
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, val)
		} else {
			n2 := stack[len(stack)-1]
			n1 := stack[len(stack)-2] // 细节: 由于栈的特性，先入后出，所以第一个计算的数字为 stack[len(stack)-2]
			// n1,n2已经出栈，取剩余的元素
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, n1+n2)
			case "-":
				stack = append(stack, n1-n2)
			case "*":
				stack = append(stack, n1*n2)
			default:
				stack = append(stack, n1/n2)
			}
		}
	}
	return stack[0]
}

// 方法二：栈顶保存最后计算结果，不过空间复杂度较高，内存占用大点
func evalRPN2(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, val)
		} else {
			n := 0
			n2 := stack[len(stack)-1]
			n1 := stack[len(stack)-2] // 细节: 由于栈的特性，先入后出，所以第一个计算的数字为 stack[len(stack)-2]
			// n1,n2已经出栈参与计算，需要把他们从栈中移除
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				n = n1 + n2
				break
			case "-":
				n = n1 - n2
				break
			case "*":
				n = n1 * n2
				break
			default:
				n = n1 / n2
			}
			stack = append(stack, n)
		}
	}
	return stack[len(stack)-1]
}
