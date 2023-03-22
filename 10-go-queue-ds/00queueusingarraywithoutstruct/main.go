//written by myself - implemented the concept of queue using array without using struct

package main

import "fmt"

var first, last int = -1, -1
var arr [10]string

func main() {
again:
	fmt.Println("Enter the operation: 1- Enqueue 2 - DeQue 3- peek")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Enter the data")
		var data string
		fmt.Scan(&data)
		enQueue(data)
		goto again

	case 2:

		deQueue()
		goto again

	case 3:

		peek()
		goto again
	case 4:
		break
	}
}

func enQueue(data string) {
	if first == 0 {
		first++
		last++
		arr[first] = data

	} else {
		first++
		arr[first] = data
	}

}

func deQueue() {
	if last < 0 {
		fmt.Println("is empty")
	} else {

		last++
		fmt.Println("process done")
	}
}

func peek() {
	if last >= 0 {
		fmt.Println(arr[last])
	} else {
		fmt.Println("is empty")
	}

}
