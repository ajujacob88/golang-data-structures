//https://www.tutorialspoint.com/go/go_recursion.htm
//https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/

package main

import "fmt"

func factorial(num int) int {
	if num <= 1 {
		fmt.Println(num)

		return 1
	}
	//fmt.Println(num)
	// fact := num * factorial(num-1)
	// return fact
	return num * factorial(num-1)
}
func main() {
	var i int = 6
	result := factorial(i)
	fmt.Println(result)
	fmt.Printf("Factorial of %d is %d", i, result)
}
