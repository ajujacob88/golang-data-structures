//https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/

// Golang program to illustrate the
// concept of indirect recursion
package main

import (
	"fmt"
)

// recursive function for
// printing all numbers
// upto the number n
func print_one(n int) {

	// if the number is positive
	// print the number and call
	// the second function
	if n >= 0 {
		fmt.Println("In first function:", n)
		// call to the second function
		// which calls this first
		// function indirectly
		print_two(n - 1)
	}
}

func print_two(n int) {

	// if the number is positive
	// print the number and call
	// the second function
	if n >= 0 {
		fmt.Println("In second function:", n)
		// call to the first function
		print_one(n - 1)
	}
}

// main function
func main() {

	// passing a positive
	// parameter which prints all
	// numbers from 1 - 10
	print_one(10)

	// this will not print
	// anything as it does not
	// follow the base case
	print_one(-1)
}
