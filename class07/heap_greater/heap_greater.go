package heap_greater

type HeapGreater struct {
	Heap       []any
	HashMap    map[any]int
	HeapSize   int
	Comparator func(a, b any) int
}

func NewHeapGreater(comparator func(a, b any) int) *HeapGreater {
	return &HeapGreater{
		Heap:       make([]any, 0),
		HashMap:    make(map[any]int),
		HeapSize:   0,
		Comparator: comparator,
	}
}

func (h *HeapGreater) IsEmpty() bool {
	return h.HeapSize == 0
}

func (h *HeapGreater) Size() int {
	return h.HeapSize
}

func (h *HeapGreater) Contains(key any) bool {
	_, ok := h.HashMap[key]
	return ok
}

func (h *HeapGreater) Peek() any {
	return h.Heap[0]
}

func (h *HeapGreater) Push(key any) {
	h.Heap = append(h.Heap, key)
	h.HashMap[key] = h.HeapSize
	h.heapInsert(h.HeapSize)
	h.HeapSize++
}

func (h *HeapGreater) Pop() any {
	ans := h.Heap[0]
	h.swap(0, h.HeapSize-1)
	delete(h.HashMap, ans)
	h.Heap = h.Heap[:h.HeapSize-1]
	h.HeapSize--
	h.heapify(0)
	return ans
}

func (h *HeapGreater) GetAll() []any {
	ans := make([]any, h.HeapSize)
	for _, val := range h.Heap {
		tmp := val
		ans = append(ans, tmp)
	}
	return ans
}

func (h *HeapGreater) Remove(key any) {
	replace := h.Heap[h.HeapSize-1]
	index := h.HashMap[key]
	delete(h.HashMap, key)
	h.Heap = h.Heap[:h.HeapSize-1]
	if key != replace {
		h.Heap[index] = replace
		h.HashMap[replace] = index
		h.Resign(replace)
	}
}

func (h *HeapGreater) Resign(key any) {
	h.heapInsert(h.HashMap[key])
	h.heapify(h.HashMap[key])
}

func (h *HeapGreater) heapInsert(index int) {
	for h.Comparator(h.Heap[index], h.Heap[(index-1)/2]) > 0 {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (h *HeapGreater) heapify(index int) {
	left := 2*index + 1
	for left < h.HeapSize {
		leftRightLargest := left
		right := left + 1
		if right < h.HeapSize && h.Comparator(h.Heap[right], h.Heap[left]) > 0 {
			leftRightLargest = right
		}
		allLargest := leftRightLargest
		if h.Comparator(h.Heap[leftRightLargest], h.Heap[index]) > 0 {
			allLargest = index
		}
		if allLargest == index {
			break
		}
		h.swap(allLargest, index)
		index = allLargest
		left = 2*index + 1
	}
}

func (h *HeapGreater) swap(i, j int) {
	o1 := h.Heap[i]
	o2 := h.Heap[j]
	h.Heap[i] = o2
	h.Heap[j] = o1
	h.HashMap[o2] = i
	h.HashMap[o1] = j
}
