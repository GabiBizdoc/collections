package heap

type HeapItem[T any] interface {
	Value() T
	Priority() int
}

type Heap[T any] []HeapItem[T]

func NewHeap[T any]() Heap[T] {
	return make(Heap[T], 0)
}

func (h *Heap[T]) Push(item HeapItem[T]) {
	*h = append(*h, item)
	h.fix(h.Len() - 1)
}

func (h *Heap[T]) Pop() HeapItem[T] {
	arr := *h
	n := (len(arr)) - 1
	if n != 0 {
		h.swap(0, n)
		if !h.down(0, n) {
			h.up(0)
		}
	}
	return h.popEnd()
}

func (h Heap[T]) Heapify() {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *Heap[T]) popEnd() HeapItem[T] {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func (h Heap[T]) fix(i int) {
	if !h.down(i, h.Len()) {
		h.up(i)
	}
}

func (h Heap[T]) Len() int {
	return len(h)
}

func (h Heap[T]) up(j int) {
	for {
		i := h.parentIndex(j)
		if i == j || !h.less(j, i) {
			break
		}
		h[j], h[i] = h[i], h[j]
		j = i
	}
}

func (h Heap[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := h.leftIndex(i)
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}

func (h Heap[T]) less(i, j int) bool {
	return h[i].Priority() > h[j].Priority()
}

func (h Heap[T]) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap[T]) leftIndex(i int) int {
	return 2*i + 1
}

func (h Heap[T]) rightIndex(i int) int {
	return 2*i + 2
}

func (h Heap[T]) parentIndex(i int) int {
	return (i - 1) / 2
}
