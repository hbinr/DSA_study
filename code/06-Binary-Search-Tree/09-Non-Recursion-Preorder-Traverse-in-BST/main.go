/*
 * @Author: duanhaobin
 * @Date: 2021-03-25 21:55:35
 */
package main

import (
	"fmt"

	"dsa.study/code/06-Binary-Search-Tree/09-Non-Recursion-Preorder-Traverse-in-BST/bst"
)

func main() {
	bst := bst.New()
	nums := [...]int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}
	bst.PreOrder()
	fmt.Println("bst: ", bst)
}
