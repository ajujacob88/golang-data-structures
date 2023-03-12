// Linked list  by Aju

//Add a node in between the nodes

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

// to Insert - Insert a value in between the nodes
func (l *linkedList) insert(index, value int) {

	// if the index greater than length, then append it
	if index > l.length {
		l.append(value)
		return //return is used instead of else, if return is not used, then use else statement..
	}

	//If the index is equal to 0, then prepend it
	if index == 0 {
		l.prepend(value)
		return
	}

	//otherwise, create a new node
	n := &node{data: value}

	//set current node to our head
	current := l.head

	//Traverse the list until we find the (index-1) node
	for i := 0; i < l.length; i++ {
		if i == index-1 {
			//set our new node next to current next
			n.next = current.next
			//set next in current value to new node
			current.next = n
			break
		}
		current = current.next
	}
	l.length++
}

func main() {

	list := linkedList{}
	list.append(20)
	list.append(30)
	list.append(40)
	list.append(50)
	list.prepend(100)
	list.append(30)
	fmt.Println("Length is:", list.len())
	list.display()
	list.insert(0, 1500)
	list.insert(30, 2500)
	list.insert(4, 6666)
	fmt.Println("Length is:", list.len())
	list.display()

}
