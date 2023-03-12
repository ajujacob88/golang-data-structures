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

//This is taken from medium - https://blog.devgenius.io/golang-data-structures-linked-list-1-2e85a90b2605#:~:text=A%20singly%20linked%20list%20is,which%20is%20the%20next%20node.&text=To%20create%20a%20singly%20linked,we%20will%20store%20our%20value.
/*
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
		Value: value,
	}
	// Check if in current list there is no node, then set the head with new node
	//else the current tail points to new node and new node becomes the tail
	if l.Length ==0{
		l.Head = el
	} else {
		l.Tail.Next = el
	}
	l.Tail = el
	l.Length +=1
}

func (l *LinkedList) prepend(value int) {
	//create new node
	el := &Node{
		Value: value,
	}
	// Check if there are nodes in our current list, if no then add new node to the tail
	// else, set the next to current head and set head to new node
	if l.Length ==0{
		l.Tail = el
	} else {
		el.Next = l.Head
	}
	l.Head = el
	l.Length +=1
}
*/

// this is good with main - https://hackthedeveloper.com/golang-linked-list/

package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

// LinkedList is a linked list
type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Len Function returns Length of the LinkedList
func (l *linkedList) Len() int {
	return l.length
}

// PushBack Function inserts a new node at the end of the LinkedList
func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l linkedList) Front() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot Find Front Value in an Empty linked list")
	}
	return l.head.data, nil
}

func (l linkedList) Back() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("Cannot Find Front Value in an Empty linked list")
	}
	return l.tail.data, nil
}

func (l *linkedList) Reverse() {
	curr := l.head
	l.tail = l.head
	var prev *node
	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	l.head = prev
}

func (l *linkedList) Delete(key int) {

	if l.head.data == key {
		l.head = l.head.next
		l.length--
		return
	}
	var prev *node = nil
	curr := l.head
	for curr != nil && curr.data != key {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		fmt.Println("Key Not found")
		return
	}
	prev.next = curr.next
	l.length--
	fmt.Println("Node Deleted")

}

func main() {
	list := linkedList{}
	node1 := &node{data: 20}
	node2 := &node{data: 30}
	node3 := &node{data: 40}
	node4 := &node{data: 50}
	node5 := &node{data: 70}
	list.PushBack(node1)
	list.PushBack(node2)
	list.PushBack(node3)
	list.PushBack(node4)
	list.PushBack(node5)
	fmt.Println("Length = ", list.Len())
	list.Display()
	list.Delete(40)
	list.Reverse()
	fmt.Println("Length = ", list.Len())
	list.Display()
	front, _ := list.Front()
	back, _ := list.Back()
	fmt.Println("Front = ", front)
	fmt.Println("Back = ", back)

}
