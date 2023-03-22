package main

import "fmt"

var top = -1
var arr [10]int

func main() {
again:
	fmt.Println("Enter the operations 1 - push, 2 - pop, 3 peek")
	var operation int
	fmt.Scanln(&operation)
	switch operation {
	case 1:
		fmt.Println("Enter the value")
		var data int
		push(data)
		goto again

	case 2:
		pop()
		goto again
	case 3:
		peek()
		goto again

	}
}

func push(data int) {
	top++

	arr[top] = data

}
func pop() {
	if top < 0 {
		fmt.Println("stack is empty")
	} else {
		top--
	}
}
func peek() {
	fmt.Println("current value is: ", arr[top])
}
