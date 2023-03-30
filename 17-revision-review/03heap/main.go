package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)

	index := len(h.array) - 1

	for h.array[parent(index)] < h.array[index] {
		h.array[parent(index)], h.array[index] = h.array[index], h.array[parent(index)]
		index = parent(index)
	}

}

func parent(i int) int {
	return (i - 1) / 2

}

func main() {
	test := &MaxHeap{}

	arr1 := []int{25, 15, 23, 65, 42, 25, 56}
	for _, v := range arr1 {
		test.insert(v)
	}

	fmt.Println(*test)

}
