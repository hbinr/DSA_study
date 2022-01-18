/*
 * @Author: duanhaobin
 * @Date: 2021-03-25 20:39:04
 */
package main

import (
	"fmt"

	"dsa.study/code/06-Binary-Search-Tree/07-InOrder-and-PostOrder-Traverse-in-BST/bst"
)

func main() {
	bst := bst.New()
	nums := [...]int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}
	bst.PostOrder()
	fmt.Println("bst: ", bst)
}
