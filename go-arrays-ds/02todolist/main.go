// Question: https://www.codechef.com/problems/TODOLIST?tab=statement
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	//defer writer.Flush()
	var testCases, noofproblems int

	fmt.Println("No of test cases:")
	fmt.Fscanln(reader, &testCases)

	for ; testCases > 0; testCases-- {
		fmt.Println("Enter the no of problems")
		fmt.Fscanln(reader, &noofproblems)
		difficultyrating := make([]int, noofproblems, noofproblems)
		count := 0
		//read the input
		fmt.Println("Enter the difficulty rating:")
		for i := 0; i < noofproblems; i++ {
			fmt.Fscan(reader, &difficultyrating[i]) //scanln will consider enter as the stop of input, thats why given fmt.Fscan only here
			if difficultyrating[i] >= 1000 {
				count++
			}
		}
		//fmt.Println("number of problems that Chef will have to remove in each case so that all remaining problems have a difficulty rating strictly less than 1000 is:")
		fmt.Fscanln(reader)
		fmt.Fprintln(writer, count)
		//writer.Flush()
	}
	fmt.Println("Number of problems that Chef will have to remove in each case so that all remaining problems have a difficulty rating strictly less than 1000 is:")
	writer.Flush()
}
