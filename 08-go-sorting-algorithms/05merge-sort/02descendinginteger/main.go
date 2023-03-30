package main

import "fmt"

func main() {

	slice1 := []int{25, 34, 12, 0, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}

	fmt.Println(slice1)

	fmt.Println("slice 1 descending by Merge sort method: ", mergeSort(slice1))

}

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
		if first[i] > second[j] {
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
