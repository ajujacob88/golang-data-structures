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
	//jenys
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
