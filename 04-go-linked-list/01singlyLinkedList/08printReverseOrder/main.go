// Linked list by Aju

//print the values in reverse order - Linked List Reverse Function

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

func (l *linkedList) reverse() {
	current := l.head
	l.tail = l.head
	var previous *node = nil
	for current != nil {
		temp := current.next
		current.next = previous
		previous = current
		current = temp
	}
	l.head = previous
}

func main() {

	list := linkedList{}
	list.append(20)
	list.append(30)
	list.append(40)
	list.append(50)
	list.prepend(100)
	list.append(30)
	list.append(30)
	list.append(30)
	list.append(30)
	list.append(30)
	fmt.Println("Length is:", list.len())
	list.display()

	fmt.Println("Reverse Order:")
	list.reverse()
	list.display()

}
