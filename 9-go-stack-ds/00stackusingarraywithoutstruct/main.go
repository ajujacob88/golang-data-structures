//written by myself - implemented the concept of stack using array without using struct
package main

import "fmt"

var top int = -1
var stack [10]int // declare array in global or pass the array while calling the function

func main() {
	// fmt.Println("Enter the size of the stack")
	// var size int
	// fmt.Scan(&size)

again:
	fmt.Println("Enter choice: 1 - Push, 2: Pop, 3: Peek, 4:Display,  6:exit")
	var choice int
	fmt.Scanln(&choice)

	for choice != 6 {
		switch choice {
		case 1:
			data := 0
			fmt.Println("Enter the data")
			fmt.Scanln(&data)
			push(data)
			goto again
		case 2:
			pop()
			goto again
		case 3:
			peek()
			goto again
		case 4:
			display()
			goto again

		case 6:
			break
		default:
			fmt.Println("Invalid choice entered")
			goto again
		}

	}
}

func push(data int) {
	top++
	stack[top] = data

}

func pop() {
	if top >= 0 {
		fmt.Println(stack[top])
		top--
	} else {
		fmt.Println("stack is empty")
	}

}
func peek() {
	if top >= 0 {
		fmt.Println(stack[top])

	} else {
		fmt.Println("stack is empty")
	}
}

func display() {
	for i := top; i >= 0; i-- {
		fmt.Println(stack[i])
	}
}
