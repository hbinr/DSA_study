package main

import "fmt"

func main() {
	var arr [10]int
	for i, arrLen := 0, len(arr); i < arrLen; i++ {
		arr[i] = i * 2
	}
	fmt.Println("arr", arr)

	scores := []int{100, 80, 60}
	for _, s := range scores {
		fmt.Println("score is ", s)
	}

	scores[1] = 90
	for _, s := range scores {
		fmt.Println("score is ", s)
	}
}
