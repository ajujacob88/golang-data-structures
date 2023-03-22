package main

import "fmt"

func main() {
	arr1 := []int{30, 11, 10, 13, 8, 10, 3, 25, 33, 1, 5, 0, 35}

	//fmt.Println(bubbleSort(arr1))

	//fmt.Println(insertionSort(arr1))

	//fmt.Println(selectionSort(arr1))

	fmt.Println(quickSort(arr1))

	//fmt.Println(mergeSort(arr1))

}

/*
func bubblesort(data []int) []int {

	for i := 0; i < len(data); i++ {
		flag := 0
		for j := 0; j < len(data)-i-1; j++ {
			if data[j+1] < data[j] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = 1
			}
		}
		if flag == 0 {
			return data
		}
	}

	return data
}

*/

/*

func insertionSort(data []int) []int {

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

*/

/*

func selectionSort(data []int) []int {

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

*/

/*

func quickSort(data []int) []int {
	lowerBount := 0
	upperBount := len(data) - 1
	return quickSortMain(data, lowerBount, upperBount)

}

func quickSortMain(data []int, lb int, ub int) []int {
	if lb < ub {
		var position int
		data, position = partition(data, lb, ub)
		data = quickSortMain(data, lb, position-1)
		data = quickSortMain(data, position+1, ub)
	}
	return data
}

func partition(data []int, lb int, ub int) ([]int, int) {
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

*/


func mergeSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	first := mergeSort(data[:len(data)/2])
	second := mergeSort(data[len(data)/2:])

	return merge(first, second)
}

func merge(first []int, second []int) []int {
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
	for ; i < len(first); i++ {
		final = append(final, first[i])
	}
	for ; j < len(second); j++ {
		final = append(final, second[j])
	}

	return final
}
*/

// once again

func quickSort(data []int) []int {

	return quickSortMain(data, 0, len(data)-1)

}

func quickSortMain(data []int, lb int, ub int) []int {
	if lb < ub {
		var pos int
		data, pos = partition(data, lb, ub)
		data = quickSortMain(data, lb, pos-1)
		data = quickSortMain(data, pos+1, ub)

	}
	return data
}

func partition(data []int, lb int, ub int) ([]int, int) {
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
