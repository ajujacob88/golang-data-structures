package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice = []string{"rahul", "aju", "eldho", "sreejith", "megha", "nikhila", "femi", "deepa", "prajeen"}

	fmt.Println(strSlice)

	fmt.Println("sorted string slice is:\n", ascendingQuickSort(strSlice))

}

func ascendingQuickSort(data []string) []string {

	return quickSort(data, 0, len(data)-1)
}

func quickSort(data []string, lb int, ub int) []string {
	var pos int
	if lb < ub {
		data, pos = partition(data, lb, ub)
		data = quickSort(data, lb, pos-1)
		data = quickSort(data, pos+1, ub)

	}

	return data
}

func partition(data []string, lb int, ub int) ([]string, int) {
	pivot := data[ub]
	i := lb

	for j := lb; j < ub; j++ {
		// if data[j] < pivot {
		// 	data[i], data[j] = data[j], data[i]
		// 	i++
		// }

		if strings.Compare(data[j], pivot) == -1 {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[ub] = data[ub], data[i]
	return data, i

}
