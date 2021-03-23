/*
 * @Author: duanhaobin
 * @Date: 2021-03-23 20:36:08
 * @LastEditTime: 2021-03-23 21:08:11
 */
package arraystack

// Stack 使用Go现有的数据结构实现栈
type Stack struct {
	element []interface{}
}

// Size 栈的大小
func (s *Stack) Size() int {
	return len(s.element)
}

// Push 入栈
func (s *Stack) Push(value ...interface{}) {
	s.element = append(s.element, value...)
}

// Peek 查看栈顶元素
func (s *Stack) Peek() interface{} {
	if s.Size() > 0 {
		return s.element[:s.Size()-1]
	}
	return nil
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	if s.Size() > 0 {
		s.element = s.element[:s.Size()-1]
	}
	return nil
}

// IsEmpty 判断栈空
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
