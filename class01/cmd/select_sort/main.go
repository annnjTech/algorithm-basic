package main

import "algorithm-basic/class01/select_sort"

func main() {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr1 := select_sort.GenerateRandomArray(maxSize, maxValue)
		arr2 := select_sort.CopyArray(arr1)
		select_sort.SelectSort(arr1)
		select_sort.Comparator(arr2)
		if !select_sort.IsEqual(arr1, arr2) {
			succeed = false
			select_sort.PrintArray(arr1)
			select_sort.PrintArray(arr2)
			break
		}
	}
	if succeed {
		print("Nice")
	} else {
		print("Fucking Fucked")
	}
}
