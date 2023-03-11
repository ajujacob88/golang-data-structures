/*
// https://www.youtube.com/watch?v=3ChyONme5WU&list=PLDtWoQ-cxqiwzMfZIVHYpyGWxLu1-mrBW&index=4
//But this is not good or not in standard form i think
package main

import (
	"fmt"
	"os"
)

type linkedList1 struct {
	number int
	next   *linkedList1
}

func main() {
	head := new(linkedList1)
	var choice int = 0
	for {
		fmt.Println("\nEnter your choice")
		fmt.Println("1 - Insert value in linked list\n2 - Display Linked List\n3 - Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var value int
			fmt.Println("Enter your value for linked list node: ")
			fmt.Scan(&value)
			head.insert(value)
		case 2:
			head.display()
		default:
			os.Exit(0)
		}

	}
}

// for insert operation
func (node *linkedList1) insert(num int) {
	var temp = &linkedList1{}
	temp.number = num
	temp.next = nil
	if node == nil {
		node = temp
	} else {
		var p = &linkedList1{}
		p = node
		for p.next != nil {
			p = p.next
		}
		p.next = temp
	}
}

//for display operation

func (node *linkedList1) display() {
	var p = &linkedList1{}
	p = node.next
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
}

*/

//This is standard form I think , taken from medium - https://blog.devgenius.io/golang-data-structures-linked-list-1-2e85a90b2605#:~:text=A%20singly%20linked%20list%20is,which%20is%20the%20next%20node.&text=To%20create%20a%20singly%20linked,we%20will%20store%20our%20value.

package main

type Node struct {
	Next   *node
	Value int
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *LinkedList) append(value int) {
	//create a new node
	el := &Node{
		Value: value
	}
}
