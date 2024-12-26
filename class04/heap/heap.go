package heap

import (
	"algorithm-basic/common"
	"errors"
)

type MyMaxHeap struct {
	HeapSize int
	Limit    int
	Heap     []int
}

func NewMyMaxHeap(limit int) *MyMaxHeap {
	return &MyMaxHeap{
		HeapSize: 0,
		Limit:    limit,
		Heap:     make([]int, limit),
	}
}

func (h *MyMaxHeap) Push(value int) (err error) {
	if h.HeapSize == h.Limit {
		return errors.New("heap is full")
	}
	h.Heap[h.HeapSize] = value
	h.heapInsert(h.HeapSize) //每一次插入都调整成成大根堆
	h.HeapSize++
	return nil
}

// 一个节点的位置是 i
// 左孩子的位置是 2*i+1
// 右孩子的位置是 2*i+2
// 父节点的位置是 (i-1)/2
func (h *MyMaxHeap) heapInsert(index int) {
	// 当前这个位置的值比父节点大的话，则交换位置
	// for停止的条件：
	// 1. 当前这个位置的值比父节点小，则停止循环
	// 2. index已经来到0位置了，则停止循环。arr[0] > arr[(0-1)/2] 肯定是false，所以停止循环
	for h.Heap[index] > h.Heap[(index-1)/2] {
		common.Swap(&h.Heap, index, (index-1)/2)
		index = (index - 1) / 2 // 然后你来到你父亲节点的位置，继续往上比较
	}
}

// Pop 弹出堆顶元素
// 返回最大值，并在大根对重，把最大值删掉
// 剩下的数，依然保持大根堆结构
func (h *MyMaxHeap) Pop() int {
	ans := h.Heap[0]                    // 先记录最大值
	h.HeapSize--                        // 堆大小-1
	common.Swap(&h.Heap, 0, h.HeapSize) // 把最后一个值放到堆顶
	h.heapify(0)                        // 调整成大根堆，从根节点开始调整，每次看自己节点和左右孩子的大小关系，然后下沉交换位置，
	return ans
}

// 从index位置，往下看，不断的下沉
// 停止的条件：
// 1.我的孩子都不在比我大
// 2.我已经没孩子了
func (h *MyMaxHeap) heapify(index int) {
	// 左孩子的位置是 2*i+1
	left := 2*index + 1
	for left < h.HeapSize { // 如果左孩子的下标不越界，说明我有孩子，如果左孩子(2*i+1)都越界了，那么右孩子(2*i+2)肯定也越界了
		/*** 先找出左右孩子的最大值 ***/
		// 假设左孩子的值最大，所以leftRightLargest默认是左孩子的位置
		leftRightLargest := left
		// 有右孩子并且右孩子的值大于左孩子的值，则leftRightLargest是右孩子的位置
		if left+1 < h.HeapSize && h.Heap[left+1] > h.Heap[left] {
			leftRightLargest = left + 1
		}
		/*** 左右孩子中最大值，跟父亲节点比较 ***/
		// 先假设父亲节点的值最大，所以largest默认是父亲节点的位置
		allLargest := index
		if h.Heap[leftRightLargest] > h.Heap[index] {
			allLargest = leftRightLargest // 左右孩子中的最大值比父亲节点大，则所有的最大值是左右孩子中的最大值的位置
		}
		if allLargest == index { // 说明两个孩子都没有父亲节点大，则不需要调整，退出循环
			break // 如果最大值是父亲节点，则不需要调整，退出循环
		}
		// 来到这里说明最大值不是父亲节点，则需要调整
		common.Swap(&h.Heap, index, allLargest) // 左右孩子中的最大值的位置跟父亲节点的位置交换
		index = allLargest                      // 当前位置index往下沉
		left = 2*index + 1                      // 重新计算当前index位置的左孩子的位置
	}
}
