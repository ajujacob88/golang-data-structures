package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice = []string{"rahul", "aju", "eldho", "sreejith", "megha", "nikhila", "femi", "deepa"}

	fmt.Println(strSlice)

	fmt.Println("sorted string slice is:\n", ascendBubbleSortOptimised(strSlice))

	// var strSlice2 = []string{"aju", "deepa", "eldho", "rahul", "megha", "nikhila", "femi", "sreejith"}

	// fmt.Println("sorted string slice is", ascendBubbleSortOptimised(strSlice2))

}

func ascendBubbleSortOptimised(data []string) []string {

	for i := 0; i < len(data); i++ {
		flag := 0
		for j := 0; j < len(data)-1; j++ {
			// if data[j+1] < data[j] {
			// 	data[j], data[j+1] = data[j+1], data[j]
			// 	flag = 1
			// }

			// //using in built function compare
			if strings.Compare(data[j+1], data[j]) == -1 {
				data[j], data[j+1] = data[j+1], data[j]
				flag = 1
			}

			//using inbuilt and for
			// for strings.Compare(data[j+1], data[j]) == -1 {
			// 	data[j], data[j+1] = data[j+1], data[j]
			// 	flag = 1
			// }
		}
		if flag == 0 {
			break
		}
		// fmt.Print(i)
	}

	return data

}
