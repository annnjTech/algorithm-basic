package select_sort

import (
	"math/rand"
	"sort"
	"time"
)

func SelectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 0 ~ n-1 找到最小值在哪，放到0位置上
	// 1 ~ n-1 找到最小值在哪，放到1位置上
	// 2 ~ n-1 找到最小值在哪，放到2位置上
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func Comparator(arr []int) {
	sort.Ints(arr)
}

func GenerateRandomArray(maxSize, maxValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, (int)(r.Float64()*float64(maxSize+1)))
	for i := 0; i < len(arr); i++ {
		// [-? , +?]
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
