package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice = []string{"rahul", "aju", "eldho", "sreejith", "megha", "nikhila", "femi", "deepa", "prajeen", "paulson"}

	fmt.Println(strSlice)

	fmt.Println("sorted string slice is:\n", mergeSort(strSlice))

}

func mergeSort(data []string) []string {
	if len(data) <= 1 {
		return data
	}
	first := mergeSort(data[:len(data)/2])
	second := mergeSort(data[len(data)/2:])
	return merge(first, second)
}
func merge(first []string, second []string) []string {
	var final []string
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		//if first[i] < second[j] {
		if strings.Compare(first[i], second[j]) == -1 {
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
