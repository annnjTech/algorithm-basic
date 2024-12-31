package main

import "algorithm-basic/class01/insert_sort"

func main() {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr1 := insert_sort.GenerateRandomArray(maxSize, maxValue)
		arr2 := insert_sort.CopyArray(arr1)
		insert_sort.InsertSort(arr1)
		insert_sort.Comparator(arr2)
		if !insert_sort.IsEqual(arr1, arr2) {
			succeed = false
			for j := 0; j < len(arr1); j++ {
				print(arr1[j], " ")
			}
			println()
			break
		}
	}
	if succeed {
		print("Nice!")
	} else {
		print("Fucking fucked!")
	}
}
