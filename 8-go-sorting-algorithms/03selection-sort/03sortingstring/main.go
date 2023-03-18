package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice = []string{"rahul", "aju", "eldho", "sreejith", "megha", "nikhila", "femi", "deepa"}

	fmt.Println(strSlice)

	fmt.Println("sorted string slice is:\n", ascendingSelectionSort(strSlice))

}

func ascendingSelectionSort(data []string) []string {
	for i := 0; i < len(data)-1; i++ {

		min := i
		//we can use the below method or inbuilt string compare package
		// for j := i + 1; j < len(data); j++ {
		// 	if data[j] < data[min] {
		// 		min = j
		// 	}
		// }

		for j := i + 1; j < len(data); j++ {
			if strings.Compare(data[j], data[min]) == -1 {
				min = j
			}
		}

		if min != i {
			data[i], data[min] = data[min], data[i]
		}

	}
	return data

}
