package heap

type Heap struct {
	Arr      []int
	HeapSize int
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) left(i int) int {
	return 2*i + 1
}

func (h *Heap) right(i int) int {
	return 2*i + 2
}

func (h *Heap) swap(i, j int) {
	h.Arr[i], h.Arr[j] = h.Arr[j], h.Arr[i]
}
