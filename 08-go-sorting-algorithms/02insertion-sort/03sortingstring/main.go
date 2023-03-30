package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice = []string{"rahul", "aju", "eldho", "sreejith", "megha", "nikhila", "femi", "deepaa"}

	fmt.Println(strSlice)

	fmt.Println("sorted string slice is:\n", ascendingInsertionSort(strSlice))

}

func ascendingInsertionSort(data []string) []string {

	//using > symbol, this also works fine

	// for i := 1; i < len(data); i++ {
	// 	temp := data[i]
	// 	j := i - 1

	// 	for ; j >= 0 && data[j] > temp; j-- {
	// 		data[j+1] = data[j]
	// 	}
	// 	data[j+1] = temp
	// }

	//using strings compare package

	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		for ; j >= 0 && strings.Compare(data[j], temp) == 1; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}
	return data
}
