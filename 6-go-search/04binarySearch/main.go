//Problem using Binary search.
// Find the index of specific character in slice

//https://dev.to/adnanbabakan/searching-algorithms-in-go-cop

package main

import "fmt"

func main() {
	var input = []int{1, 2, 3, 4, 5, 6, 7, 8}

	number := FindIndex(input, 7)

	fmt.Println(number)
}

func FindIndex(arr []int, searchValue int) int {
	first := 0
	last := len(arr) - 1
	//fmt.Println(high)
	for first <= last {
		mid := int((first + last) / 2)
		//fmt.Println(mid)
		if arr[mid] == searchValue {
			return mid
		} else if arr[mid] < searchValue {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return -1

}
