//sorting a slice in ascending order using bubble sort
// https://blog.devgenius.io/sorting-algorithm-bubble-selection-vs-insertion-vs-merge-which-one-is-the-best-c30ea1a58629

package main

import "fmt"

func main() {

	var slice1 = []int{25, 12, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}
	fmt.Println(slice1)

	// slice2 := ascendSort(slice1)
	// fmt.Println(slice1) //since slice is mutable, both slice1 and slice2 will be replaced
	// fmt.Println(slice2)

	// ascendSort(slice1)
	// fmt.Println(slice1)

	fmt.Println("\nslice 1 sorted: ", ascendSortBubbleSort(slice1))

	var slice3 = []int{2, 3, 1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println(slice3)

	fmt.Println("\nslice 1 sorted: ", optimisedBubbleSort(slice3)) //since we used optimised sorting the i for loop run here only 2 times, otherwise it will run 14 times, so optimised is good...

}

func ascendSortBubbleSort(data []int) []int {

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ { //since -i-1 because no need to go to end after each iteration
			if data[j+1] < data[j] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
		fmt.Print(i, " , ") //no of times  for loop i worked
	}
	return data
}

func optimisedBubbleSort(data []int) []int { //ref: https://medium.com/@michellekwong2/what-to-know-about-bubble-sort-f0f392905008

	for i := 0; i < len(data); i++ {
		flag := 0
		for j := 0; j < len(data)-i-1; j++ {
			if data[j+1] < data[j] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
		fmt.Print(i, " , ") //no of times  for loop i worked
	}
	return data

}

/*
Bubble sort has a time complexity of O(n²), which is not really performant.
The reason is that we have a nested loop to compare each element.
So, if we have 5 elements, we would have 25 operations to sort a list.
For the space complexity, bubble sort has O(1) which is very performant since we stored the result in the same array.



Also, the best case time complexity will be O(n), it is when the list is already sorted. it may happen if use optimisedBubbleSort function ..
The best time complexity for Bubble Sort is O(n). The average and worst time complexity is O(n²).
*/
