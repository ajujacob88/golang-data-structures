// https://www.codechef.com/problems/PRACTICEPERF?tab=statement
package main

import "fmt"

func main() {
	var p [4]int
	fmt.Println("Enter the no of problems solved for each weeks")
	count := 0
	for i := 0; i < 4; i++ {
		fmt.Scan(&p[i])
		if p[i] >= 10 {
			count++
		}
	}
	fmt.Println("The number of weeks in which Chef solved at least 10 problems is: ", count)

}
