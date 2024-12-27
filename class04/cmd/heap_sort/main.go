package main

import (
	"algorithm-basic/class04/heap"
	"fmt"
)

func main() {

	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	heap.HeapSort(arr)
	fmt.Println(arr)
}
