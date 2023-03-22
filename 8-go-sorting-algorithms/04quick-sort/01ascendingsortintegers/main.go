// https://www.youtube.com/watch?v=QN9hnmAgmOc&list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU&index=101
//https://blog.boot.dev/golang/quick-sort-golang/

package main

import "fmt"

func main() {

	slice1 := []int{25, 34, 12, 0, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}

	fmt.Println(slice1)

	fmt.Println("slice 1 ascending by Quick sort method: ", ascendingQuickSort(slice1))

}

func ascendingQuickSort(data []int) []int {
	lowerBount := 0
	upperBount := len(data) - 1
	return quickSort(data, lowerBount, upperBount)

}

func quickSort(data []int, lb int, ub int) []int {
	if lb < ub {
		var position int
		data, position = partition(data, lb, ub)
		data = quickSort(data, lb, position-1)
		data = quickSort(data, position+1, ub)

	}
	//fmt.Println(data)
	return data
}

func partition(data []int, lb int, ub int) ([]int, int) {
	//jenys no need, study other blog.boot.dev
	// pivot := data[lb]
	// start := lb
	// end := ub

	// for start < end {
	// 	for data[start] <= pivot {
	// 		start++
	// 	}
	// 	for data[end] > pivot {
	// 		end--
	// 	}
	// 	if start < end {
	// 		data[start], data[end] = data[end], data[start]
	// 	}
	// }
	// data[lb], data[end] = data[end], data[lb]

	// return data, end

	//blog.boot.dev - this is good

	pivot := data[ub]
	i := lb
	for j := lb; j < ub; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}

	data[i], data[ub] = data[ub], data[i]

	return data, i

}

/*

Time Complexity of Quick Sort
In Worst Case - O(n²) is the time complexity of selection sort
In average case or best case - O(nlogn) which is far better than the bubble, selection and insertion sort..


On average, quicksort has a Big O of O(n*log(n)).
In the worst case, and assuming we don’t take any steps to protect ourselves, it can break down to O(n^2).
The partition() function has a single for-loop that ranges from the lowest index to the highest index in the array.
By itself, the partition() function is O(n).
The overall complexity of quicksort is dependent on how many times partition() is called.

In the worst case, the input is already sorted.
An already sorted array results in the pivot being the largest or smallest element in the partition each time.
When this is the case, partition() is called a total of n times.
In the best case, the pivot is the middle element of each sublist which results in log(n) calls to partition().


Finding the median optimization
One of the most popular solutions is to use the “median of three” approach.
Three elements (for example: the first, middle, and last elements) of each partition are chosen and the median is found between them.
That item is then used as the pivot.
This approach has the advantage that it can’t break down to O(n^2) time because we are guaranteed to never use the worst item in the partition as the pivot.
That said, it can still be slow_er_ because a true median isn’t used.
*/
