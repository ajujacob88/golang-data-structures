//https://www.codechef.com/problems/FCTRL2

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var cases, num int64
	fmt.Println("No of test cases")
	fmt.Scan(&cases)
	for cases != 0 {
		fmt.Print("Enter the number: ")
		fmt.Scan(&num)
		var fact = new(big.Int)
		fact.MulRange(1, num)
		fmt.Println("Factorial of ", num, " is ", fact)
		cases--
	}

}

/*
You are asked to calculate factorials of some small positive integers.

Input
An integer t, 1<=t<=100, denoting the number of testcases, followed by t lines, each containing a single integer n, 1<=n<=100.

Output
For each integer n given at input, display a line with the value of n!

*/
