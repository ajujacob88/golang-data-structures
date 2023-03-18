package main

import "fmt"

func main() {

	var slice1 = []int{25, 12, 52, 9, 2, 8, 6, 64, 48, 25, 36, 12, 2, 1}
	fmt.Println(slice1)

	fmt.Println("slice 1 descending sorted: ", descendingInsertionSort(slice1))

}

func descendingInsertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		for ; j >= 0 && data[j] < temp; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}
	return data
}
