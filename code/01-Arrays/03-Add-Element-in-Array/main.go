package main

import (
	"fmt"

	"dsa.study/code/01-Arrays/03-Add-Element-in-Array/array"
)

func main() {
	arr := array.NewArray(5)

	arr.AddLast(30)
	arr.AddFirst(10)
	arr.AddFirst(20)
	arr.AddFirst(40)
	fmt.Println("array is :", arr)
}
