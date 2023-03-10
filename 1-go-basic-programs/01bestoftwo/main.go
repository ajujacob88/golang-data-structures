//https://www.codechef.com/problems/BESTOFTWO

package main

import (
	"fmt"
	"math"
)

func main() {
	var testCases int
	var x, y float64
	fmt.Println("Enter the no of test cases")
	fmt.Scan(&testCases)
	fmt.Println("Marks scored in first attempt and second attempt")
	for i := 0; i < testCases; i++ {
		fmt.Scan(&x, &y)
		fmt.Print("The max score is ")
		fmt.Println(math.Max(x, y))
	}

}

/*
Problem
Chef took an examination two times. In the first attempt, he scored X marks while in the second attempt he scored Y marks.
According to the rules of the examination, the best score out of the two attempts will be considered as the final score.
Determine the final score of the Chef.

Input Format
The first line contains a single integer T — the number of test cases. Then the test cases follow.
The first line of each test case contains two integers X and Y — the marks scored by Chef in the first attempt and second attempt respectively.

Output Format
For each test case, output the final score of Chef in the examination.

Constraints
1≤T≤1000
0≤X,Y≤100
*/
