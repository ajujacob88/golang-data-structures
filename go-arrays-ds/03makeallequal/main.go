//https://www.codechef.com/problems/PAIREQ

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	write := bufio.NewWriter(os.Stdout)
	defer write.Flush()

	var testCases, size int
	fmt.Fscan(read, &testCases)
	for i := 0; i < testCases; i++ {
		m := make(map[int]int)
		fmt.Fscan(read, &size)

		var value int
		for j := 0; j < size; j++ {
			fmt.Fscan(read, &value)
			m[value] = m[value] + 1
			//fmt.Printf("index is %v and value is %v \n", value, m[value])
		}
		min := math.MaxInt32 //its a constant value maximum of int32 value will be stored in min,,, we can replace this with a large no also.
		//min := 999999999   //this will also works fine instead of above
		for _, count := range m {
			if size-count < min {
				min = size - count
			}
		}
		fmt.Fprintln(write, min)
	}

}
