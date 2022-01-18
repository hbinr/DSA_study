/*
 * @Author: duanhaobin
 * @Date: 2021-03-23 15:26:04
 * @LastEditTime: 2021-03-23 15:29:27
 */
package main

import (
	"fmt"

	"dsa.study/code/04-Linked-List/05-Remove-Element-in-LinkedList/linkedlist"
)

func main() {
	linkedList := linkedlist.New()
	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
	}
	fmt.Println(linkedList)

	linkedList.Remove(1)
	fmt.Println(linkedList)

	linkedList.RemoveFirst()
	fmt.Println(linkedList)

	linkedList.RemoveLast()
	fmt.Println(linkedList)
}
