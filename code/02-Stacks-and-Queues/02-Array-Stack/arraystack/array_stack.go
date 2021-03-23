/*
 * @Author: duanhaobin
 * @Date: 2021-03-23 18:22:34
 * @LastEditTime: 2021-03-23 18:38:30
 */

package arraystack

import (
	"fmt"
	"strings"

	"dsa.study/code/02-Stacks-and-Queues/02-Array-Stack/array"
)

type ArrayStack struct {
	array *array.Array
}

func New(capacity int) *ArrayStack {
	return &ArrayStack{
		array: array.New(capacity),
	}
}

func (as *ArrayStack) GetSize() int {
	return as.GetSize()
}

func (as *ArrayStack) IsEmpty() bool {
	return as.IsEmpty()
}

// 压入栈顶元素
func (as *ArrayStack) Push(e interface{}) {
	as.array.AddLast(e)
}

// 弹出栈顶元素
func (as *ArrayStack) Pop() interface{} {
	return as.array.RemoveLast()
}

// 查看栈顶元素
func (as *ArrayStack) Peek() interface{} {
	return as.array.GetLast()
}

func (as *ArrayStack) String() string {
	var builder strings.Builder
	builder.WriteString("Stack: ")
	builder.WriteString("[")
	for i, count := 0, as.array.GetSize(); i < count; i++ {
		builder.WriteString(fmt.Sprint(as.array.Get(i)))
		// 中间用 ',' 隔开
		if i != as.array.GetSize()-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}
