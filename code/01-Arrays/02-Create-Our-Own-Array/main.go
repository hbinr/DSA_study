package main

import (
	"fmt"

	"dsa.study/code/01-Arrays/02-Create-Our-Own-Array/array"
)

func main() {
	arr := array.NewArray(20)
	fmt.Println(arr.GetSize())
	fmt.Println(arr.GetCapacity())
	fmt.Println(arr.IsEmpty())
}
