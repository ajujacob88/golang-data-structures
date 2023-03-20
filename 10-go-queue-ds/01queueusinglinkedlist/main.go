// https://blog.devgenius.io/golang-data-structures-stack-and-queue-2-5a6822baeb0d

package main

import "fmt"

type node struct {
	data string
	next *node
}

type queues struct {
	length      int
	first, last *node
}

func main() {

	var choice int
	que := queues{}

again:
	fmt.Println("Enter choice: 1 - enQue, 2: deQue, 3: Peek, 4: Display")
	fmt.Scanln(&choice)

	for choice != 6 {
		switch choice {
		case 1:
			var data string
			fmt.Println("Enter the data")
			fmt.Scanln(&data)
			que.enQue(data)
			goto again
		case 2:
			que.deQue()
			goto again
		case 3:
			que.peek()
			goto again

		case 4:
			que.display()
			goto again
		default:
			fmt.Println("Invalid choice entered")
			goto again
		}
	}
}

func (q *queues) enQue(value string) {
	n := &node{data: value}

	if q.length == 0 {
		q.first = n
	} else {
		q.last.next = n
	}
	q.last = n
	q.length++
}

func (q *queues) deQue() {
	if q.length <= 0 {
		fmt.Println("Queue is Empty")
		return
	}
	if q.length == 1 {
		q.last = nil
	}
	first := q.first
	q.first = first.next
	first.next = nil
	q.length--

}

func (q *queues) peek() {
	fmt.Println(q.first)
}

func (q *queues) display() {
	currentNode := q.first
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

/*

Stack and queue are one of the data structures that software engineers need to know.
They can be implemented in a lot of study cases. Stack is helpful when you want to get things out in reverse order from when you put them in.
One example of a stack is when the computer reads our program.
Queue is helpful when things don’t have to be processed immediately, but have to be processed in First In First Out order.
One example of a queue is when we buy tickets at the cinema.

*/
