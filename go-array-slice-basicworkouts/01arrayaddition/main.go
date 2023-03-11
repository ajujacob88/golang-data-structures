package main

import "fmt"

func main() {
	var array1 = [10]int{1, 2, 3, 4, 5}
	var array2 = [10]int{2, 4, 6, 8, 10}
	var sum [10]int
	array1[0] = 10

	fmt.Println(array1)
	fmt.Println(array2)
	for i := 0; i < 10; i++ {
		sum[i] = array1[i] + array2[i]
	}
	fmt.Println(sum)

}
