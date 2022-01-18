/*
 * @Author: duanhaobin
 * @Date: 2021-03-23 18:37:04
 */
package main

import (
	"fmt"

	"dsa.study/code/02-Stacks-and-Queues/02-Array-Stack/arraystack"
)

func main() {
	stack := arraystack.New(10)
	fmt.Println("stack: ", stack)

	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	fmt.Println(stack)

	fmt.Println("peek(): ", stack.Peek())
	fmt.Println("pop(): ", stack.Pop())
	fmt.Println(stack)

}
