package main

import "fmt"

func main() {
	origSlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice: ", origSlice)
	fmt.Println("Enter the index to delete: ")
	var removepos int
	fmt.Scan(&removepos)
	newSlice := append(origSlice[:removepos], origSlice[removepos+1:]...)

	fmt.Println("New Slice:", newSlice)

}

//https://www.geeksforgeeks.org/delete-elements-in-a-slice-in-golang/
