package main

import "fmt"

func main() {
	originalArray := [5]int{100, 200, 300, 400, 500}
	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Enter the index to delete: ")
	var removepos int
	fmt.Scan(&removepos)
	// check if the index is within array bounds
	if removepos < 0 || removepos >= len(originalArray) {
		fmt.Println("Index out of rang")
	} else {
		// delete an element from the array
		newlength := 0
		for index := range originalArray {
			if removepos != index {
				originalArray[newlength] = originalArray[index]
				newlength++
			}
		}
		// reslice the array to remove extra index
		newArray := originalArray[:newlength]
		fmt.Println("The new array is: ", newArray)
	}
	//fmt.Println(originalArray)

}
