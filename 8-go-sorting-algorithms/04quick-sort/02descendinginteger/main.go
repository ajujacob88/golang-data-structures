package main

import "fmt"

func main() {

	slice1 := []int{25, 34, 12, 0, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}

	fmt.Println(slice1)

	fmt.Println("slice 1 descending by Quick sort method: ", descendingQuickSort(slice1))

}

func descendingQuickSort(data []int) []int {

	return quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, lb int, ub int) []int {
	var pos int
	if lb < ub {
		data, pos = partition(data, lb, ub)
		data = quickSort(data, lb, pos-1)
		data = quickSort(data, pos+1, ub)

	}

	return data
}

func partition(data []int, lb int, ub int) ([]int, int) {
	pivot := data[ub]
	i := lb

	for j := lb; j < ub; j++ {
		if data[j] > pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[ub] = data[ub], data[i]
	return data, i

}
