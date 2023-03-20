// https://www.youtube.com/watch?v=jlHkDBEumP0&list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU&index=102
// https://blog.boot.dev/golang/merge-sort-golang/

package main

import "fmt"

func main() {

	slice1 := []int{25, 34, 12, 0, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}

	fmt.Println(slice1)

	fmt.Println("slice 1 ascending by Merge sort method: ", mergeSort(slice1))

}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		//fmt.Println(data)
		return data
	}
	first := mergeSort(data[:len(data)/2])
	second := mergeSort(data[len(data)/2:])
	return merge(first, second)

}

func merge(first []int, second []int) []int {
	//fmt.Println("first", first, "second", second)
	final := []int{}
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}
	//fmt.Println("final is", final)
	for ; i < len(first); i++ {
		final = append(final, first[i])
	}
	//fmt.Println("final is", final)
	for ; j < len(second); j++ {
		final = append(final, second[j])
	}
	//fmt.Println("final is", final)
	return final
}

/*

Time Complexity of Merge Sort
Both in Best Case and Worst Case -O(nlogn) is the time complexity of merge sort

Pros
Fast. Merge sort is much faster than bubble sort, being O(n*log(n)) instead of O(n^2).
Stable. Merge sort is also a stable sort which means that values with duplicate keys in the original list will be in the same order in the sorted list.
Cons
Extra memory. Most sorting algorithms can be performed using a single copy of the original array.
Merge sort requires an extra array in memory to merge the sorted subarrays.
Recursive: Merge sort requires many recursive function calls, and function calls can have significant resource overhead.

*/
