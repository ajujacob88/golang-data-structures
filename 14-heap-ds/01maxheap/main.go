//https://www.youtube.com/watch?v=3DYIgTC4T1o
//https://www.youtube.com/watch?v=NEtwJASLU8Q&list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU&index=104
// https://maneeshaindrachapa.medium.com/heap-data-structure-in-golang-98641a32d2e3

package main

import "fmt"

//MaxHeap struct has a slice that holds the array - (accessibl from outside, so should start with capital)
type MaxHeap struct {
	array []int
}

//insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract or Delete returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	//when the array is empty
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}

	//take out the last index and put it int0 the root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	//loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { //when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { //when left child is larger
			childToCompare = l
		} else { //when right child is larger
			childToCompare = r
		}
		//compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

//get the parent index
func parent(i int) int {
	return (i - 1) / 2

}

//get the left child index
func left(i int) int {
	return 2*i + 1
}

//get the right child index
func right(i int) int {
	return 2*i + 2
}

//swap keys/values in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	//fmt.Println(m)
	//buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	buildHeap := []int{50, 45, 35, 33, 16, 25, 34, 12, 10, 120}
	for _, v := range buildHeap {
		m.Insert(v)
		//fmt.Println(m)
	}
	fmt.Println("The heap is ", *m)
	fmt.Println("The first extracted value is", m.Extract())
	fmt.Println(m)

	for i := 0; i < 9; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
