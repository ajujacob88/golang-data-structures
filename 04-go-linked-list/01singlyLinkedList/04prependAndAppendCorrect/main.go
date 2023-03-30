// Linked list final correct one by Aju - this is the final one

//Add a node at the end & beginning

//https://blog.devgenius.io/golang-data-structures-linked-list-1-2e85a90b2605

// https://hackthedeveloper.com/golang-linked-list/

package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// to find length
func (l *linkedList) len() int {
	return l.length
}

// to display
func (l *linkedList) display() {
	temp := l.head //inorder to store the address of head, otherwise head address will be lost and if we print again after two steps ,then the initial head address is lost and not able to print the values.. this code written by myself
	for l.head != nil {
		fmt.Printf("%v ->", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
	l.head = temp
}

// to append - means add data at the end of the linked  list
func (l *linkedList) append(value int) {
	n := &node{data: value}
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

// to prepend - means add data to the beginning of the linked  list
func (l *linkedList) prepend(value int) {
	n := &node{data: value}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		temp := l.head
		n.next = temp
		l.head = n
		l.length++
	}

}

func main() {

	list := linkedList{}
	list.append(20)
	list.append(30)
	list.append(40)
	list.append(50)
	list.prepend(100)
	list.append(30)
	// fmt.Println("Length is:", list.len())
	// list.display()
	list.append(5)
	list.prepend(200)
	list.append(3)
	// fmt.Println("Length is:", list.len())
	// list.display()

	list.display()

}
