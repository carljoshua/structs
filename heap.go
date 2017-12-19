package structs

import (
    "math"
    "fmt"
)

type Heap struct {
    data        []int
    num_nodes   int
}

// NewHeap() returns a heap
func NewHeap() *Heap {
    return &Heap{}
}

// Insert() stores any number of parameters into the heap
func (h *Heap) Insert(v ...int) {
    for _, val := range(v){
        h.data = append(h.data, val)
    }

    n := len(h.data)

    i := (float64) ((n / 2) - 1)
    h.num_nodes = (int) (math.Floor(i))
}

// MaxHeap() converts the heap into a max-heap
func (h *Heap) MaxHeap() {
    for i := h.num_nodes; i >= 0; i-- {
        if h.data[i] < h.data[i * 2 + 1] {
            swap(h.data, i, i * 2 + 1)
        }else if len(h.data) > i * 2 + 2 && h.data[i] < h.data[i * 2 + 2] {
            swap(h.data, i, i * 2 + 2)
        }
    }

    if !h.MaxHeapTest() {
        h.MaxHeap()
    }
}

// MaxHeapTest() evaluate the heap if it is a max-heap
func (h *Heap) MaxHeapTest() bool {
    ok := true
    for i := h.num_nodes; i >= 0; i-- {
        node := h.data[i]
        if node < h.data[i * 2 + 1] || (len(h.data) > i * 2 + 2 && node < h.data[i * 2 + 2]) {
            ok = false
            break
        }
    }
    return ok
}

// MinHeap() converts the heap into a min-heap
func (h *Heap) MinHeap() {
    for i := h.num_nodes; i >= 0; i-- {
        if h.data[i] > h.data[i * 2 + 1] {
            swap(h.data, i, i * 2 + 1)
        }else if len(h.data) > i * 2 + 2 && h.data[i] > h.data[i * 2 + 2] {
            swap(h.data, i, i * 2 + 2)
        }
    }

    if !h.MinHeapTest() {
        h.MinHeap()
    }
}

// MinHeapTest() evaluate the heap if it is a min-heap
func (h *Heap) MinHeapTest() bool {
    ok := true
    for i := h.num_nodes; i >= 0; i-- {
        node := h.data[i]
        if node > h.data[i * 2 + 1] || (len(h.data) > i * 2 + 2 && node > h.data[i * 2 + 2]) {
            ok = false
            break
        }
    }
    return ok
}

// Print() prints the data in the heap into the console
func (h *Heap) Print() {
    fmt.Println(h.data)
}

// swap() swaps the elements in the given indexes in the given array
func swap(data []int, i int, j int) {
    temp := data[i]
    data[i] = data[j]
    data[j] = temp
}
