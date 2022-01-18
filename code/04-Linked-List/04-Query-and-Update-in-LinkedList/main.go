/*
 * @Author: duanhaobin
 * @Date: 2021-03-23 12:59:45
 * @LastEditTime: 2021-03-23 14:57:33
 */
package main

import (
	"fmt"

	"dsa.study/code/04-Linked-List/04-Query-and-Update-in-LinkedList/linkedlist"
)

func main() {
	linkedList := linkedlist.New()
	for i := 0; i < 1; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}
	linkedList.AddLast(100)
	fmt.Println(linkedList)
}
