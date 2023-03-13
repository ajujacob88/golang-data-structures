//https://www.tutorialspoint.com/go/go_recursion.htm
//https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/

package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
}

func fibonaci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonaci(n-1) + fibonaci(n-2)
}
