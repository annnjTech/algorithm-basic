package heap

import "algorithm-basic/common"

// HeapSort 堆排序
// 时间复杂度：O(N*logN)
// 空间复杂度：O(1)
func HeapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	// 经典，建堆1，O(N*logN)
	for i := 0; i < len(arr); i++ { // O(N)
		heapInsert(arr, i) // O(logN)
	}
	// 优化，建堆2，O(N)
	// T(N) =
	for i := len(arr) - 1; i > 0; i-- { // O(N)
		heapify(arr, i, len(arr))
	}
	heapSize := len(arr)
	common.Swap(&arr, 0, heapSize-1) // 交换堆顶和最后一个元素
	heapSize--                       // 堆大小减1
	for heapSize > 0 {               // 堆不为空
		heapify(arr, 0, heapSize) // 向下调整堆
		common.Swap(&arr, 0, heapSize-1)
		heapSize--
	}

}
