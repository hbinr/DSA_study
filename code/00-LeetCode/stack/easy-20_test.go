/*
 * @Author: duanhaobin
 * @Date: 2021-03-23 19:41:37
 */
package stack

import (
	"fmt"
	"testing"
)

// 方法1思路:
// 1.先创建map来保存括号对应关系，key为反括号，value为对应key匹配括号类型。注意细节，key值为反括号
// 2.遍历入参字符串长度，然后从map中取一个括号入栈（随机的），此时该括号元素为栈顶
// 3pairs[s[i]] > 0
// 3.判断栈顶元素(stack[len(stack)-1])和字符串中给定的某个括号(pairs[s[i]])是否相同：
//   补充：
//  	 - (stack[len(stack)-1]) 的值为入参字符串中的某个括号
//   	 - (pairs[s[i]])  只有三个值： (, [, {
//    3.1 不相同，如 '(' != ']' ,显然这不满足括号匹配的条件，直接返回false
//    3.2 相同，如 '(' == '('，则后续还需要再判断是否有一个字符为 ‘)’ 才能满足调件。
//    这儿有个技巧，如果相同，直接将栈顶元素出栈，再看下一个括号
//    这样就保证栈顶始终只维持是一个元素，
//

func TestEasy20(t *testing.T) {
	fmt.Println("res: ", isValid1("(){}}"))
	fmt.Println("res: ", isValid2("(){}}"))
}

func isValid1(s string) bool {
	// 1.只有偶数个字符才能匹配成功，奇数直接返回false
	n := len(s)
	if n%2 == 1 {
		return false
	}

	// 2.创建map来保存括号对应关系，key为反括号，value为对应key匹配括号类型。注意细节，key值为反括号
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 3.以栈的方式来保存括号
	stack := []byte{}
	for i := 0; i < n; i++ {
		// 4. pairs[s[i]] > 0时 , 只有三种情况： (, [, {，那么此时s[i] 为 ), ], }
		if pairs[s[i]] > 0 {
			// 6. 栈非空时，如果栈顶元素 != (, [, { 这三种情况之一，则不匹配
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 7. 如果栈顶元素 == (, [, { ，那么意味着 s[i]为 ), ], }之一，即满足了括号匹配，栈顶出栈，直到栈空
			stack = stack[:len(stack)-1]
		} else {
			// 5. 栈顶元素只保存左括号类型 s[i]为 (, [, {
			stack = append(stack, s[i])
		}
	}

	// 8. 栈空表示括号都匹配
	return len(stack) == 0
}

//  方法二；更容易理解
func isValid2(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[rune]rune{ // rune 是因为string底层为rune（type int32），更底层为Unicode编码，最多4个字节，
		// 再更底层为utf-8编码
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}
	for _, char := range s {
		// 1. 将左括号类型全部入栈
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)

		} else if len(stack) != 0 && stack[len(stack)-1] == pairs[char] { // 2.如果括号匹配，则出战
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
