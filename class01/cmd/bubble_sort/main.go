package main

import "algorithm-basic/class01/bubble_sort"

func main() {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr1 := bubble_sort.GenerateRandomArray(maxSize, maxValue)
		arr2 := bubble_sort.CopyArray(arr1)
		bubble_sort.BubbleSort(arr1)
		bubble_sort.Comparator(arr2)
		if !bubble_sort.IsEqual(arr1, arr2) {
			succeed = false
			break
		}
	}
	if succeed {
		print("Nice!")
	} else {
		print("Fucking Fucked!")
	}
}
