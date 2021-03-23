/*
 * @Author: duanhaobin
 * @Date: 2021-03-23 17:00:49
 * @LastEditTime: 2021-03-23 17:27:32
 */
// Stack 栈基础操作接口
package stack

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
}
