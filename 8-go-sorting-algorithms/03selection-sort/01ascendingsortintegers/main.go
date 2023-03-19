//https://www.youtube.com/watch?v=9oWd4VJOwr0&list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU&index=99
//Selection Sort

package main

import "fmt"

func main() {

	slice1 := []int{25, 12, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}

	fmt.Println(slice1)

	fmt.Println("slice 1 ascending by selection sorted: ", ascendingSelectionSort(slice1))
}

func ascendingSelectionSort(data []int) []int {

	for i := 0; i < len(data)-1; i++ {

		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		if min != i {
			data[i], data[min] = data[min], data[i]
		}

	}
	return data
}

/*
Time Complexity
Both in Best Case and Worst Case - O(n²) is the time complexity of selection sort


Space Complexity
Space Complexity of selection sort is O(1)

*/
