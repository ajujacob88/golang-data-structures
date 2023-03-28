// heap sort in maxheap will gives us the array in ascending order..

// https://www.youtube.com/watch?v=NEtwJASLU8Q&list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU&index=103   - last part of this video
//https://www.youtube.com/watch?v=Q_eia3jC9Ts&list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU&index=104

// by sethukumar
package main

import "fmt"

func HeapSort(array []int, endIdx int) []int {
	if endIdx <= 0 {
		return array
	}
	sorted := Heapify(array, endIdx)
	swap(sorted, endIdx, 0)
	return HeapSort(sorted, endIdx-1)

}

func Heapify(array []int, endIdx int) []int {

	for i := parent(endIdx); i >= 0; i-- {
		ShiftDown(array, endIdx, i)
	}

	return array
}

func ShiftDown(array []int, endIdx, currentIdx int) []int {

	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {

		rightIdx := rightChild(currentIdx)
		var idxToShift int
		if rightIdx <= endIdx && array[rightIdx] > array[leftIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}

		if array[idxToShift] > array[currentIdx] {

			swap(array, idxToShift, currentIdx)
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			return array
		}
	}

	return array
}

func main() {
	array := []int{5, 6, 3, 6, 3, 2, 17}

	sortedArray := HeapSort(array, len(array)-1)

	fmt.Println("Sorted array:", sortedArray)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func leftChild(i int) int {
	return (i * 2) + 1
}
func rightChild(i int) int {
	return (i * 2) + 2
}
func parent(i int) int {
	return (i - 1) / 2
}

/*
Time complexity
for insertion of 1 element - o(logn)
but here, all elements(n) should be inserted because, building heap array, so time complexity would be o(nlogn)

llly for deleting 1 element time complexity is o(logn)
for n elements, it is o(nlogn).. here n elements need to be deleted..

so total time complexity of heapsort would be o(2nlogn) which is equal to o(nlogn) , both in worst and best and every case...  (since 2 is constant and can be omited)
*/

/*
//anothermethod but not correct method i think,
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

//not correct method i think, because new array is created
func (h *MaxHeap) heapSortAscend() {

	sortedNums := []int{}
	for len(h.array) > 0 {
		sortedNums = append(sortedNums, h.Extract())
	}
	fmt.Println("Ascending sorted values: ", sortedNums)

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
	// fmt.Println("The first extracted value is", m.Extract())
	// fmt.Println(m)

	// for i := 0; i < 9; i++ {
	// 	m.Extract()
	// 	fmt.Println(m)
	// }

	//m.heapSortAscend()
	// m2 := &MaxHeap{}
	// array2 := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	m.heapSortAscend()

}


*/
