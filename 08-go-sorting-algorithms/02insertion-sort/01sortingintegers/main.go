//https://www.youtube.com/watch?v=yCxV0kBpA6M&list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU&index=99

package main

import "fmt"

func main() {

	var slice1 = []int{25, 12, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}
	fmt.Println(slice1)

	fmt.Println("slice 1 ascending sorted: ", ascendingInsertionSort(slice1))

}

func ascendingInsertionSort(data []int) []int {

	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > temp; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}

	return data

}

/*
Time Complexity
Worst Case- O(n²)
Best Case- O(n) – When the array is already sorted

Space Complexity
Space Complexity of insertion sort is O(1)

Insertion sort is used when the numbers are pretty sure sorted or we have small datasets.
Insertion sort is much less efficient on large lists than more advanced algorithms like quicksort or merge sort.
Insertion sort is a simple algorithm that works just like you would arrange playing cards in your hands.
A slice is first split into sorted and unsorted sections, then values from the unsorted section are inserted into the correct position in the sorted section.

*/
