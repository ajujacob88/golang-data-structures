// Linked list by Aju

//Delete a node based on the value

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

// to Delete - delete a value in between the nodes
func (l *linkedList) delete(value int) {

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	var previous *node = nil
	current := l.head

	for current != nil && current.data != value {
		previous = current
		current = current.next
	}
	if current == nil {
		fmt.Println("Value not found")
		return
	}
	previous.next = current.next
	l.length--
	fmt.Println("Node Deleted")
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

	list.delete(100)
	fmt.Println("Length is:", list.len())
	list.display()
	list.delete(50)
	fmt.Println("Length is:", list.len())
	list.display()
	list.delete(525)
	fmt.Println("Length is:", list.len())
	list.display()
	list.delete(30)
	fmt.Println("Length is:", list.len())
	list.display()

}
