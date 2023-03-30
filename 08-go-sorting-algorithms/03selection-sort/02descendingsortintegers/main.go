package main

import "fmt"

func main() {

	slice1 := []int{25, 12, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}

	fmt.Println(slice1)

	fmt.Println("slice 1 descending by selection sorted: ", descendingSelectionSort(slice1))
}

func descendingSelectionSort(data []int) []int {

	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] > data[min] {
				min = j
			}
		}
		if min != i {
			data[i], data[min] = data[min], data[i]
		}
	}
	return data
}
