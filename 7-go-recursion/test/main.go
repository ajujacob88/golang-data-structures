package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonaci(i))
	}
}

func fibonaci(num int) int {
	if num == 0 {
		return 1
	}
	if num == 1 {
		return 1
	}

	return fibonaci(num-1) + fibonaci(num-2)

}
