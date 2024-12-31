package bubble_sort

import (
	"math/rand"
	"sort"
	"time"
)

func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 0 ~ n-1
	// 0 ~ n-2
	// 0 ~ n-3
	for e := len(arr) - 1; e > 0; e-- { // 0 ~ e
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func Comparator(arr []int) {
	sort.Ints(arr)
}

func GenerateRandomArray(maxSize, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, (int)(r.Float64()*float64(maxSize+1)))
	for i := 0; i < len(arr); i++ {
		arr[i] = (int)(r.Float64()*float64(maxValue+1)) - (int)(r.Float64()*float64(maxValue))
	}
	return arr
}

func CopyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func IsEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func PrintArray(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
	println()
}
