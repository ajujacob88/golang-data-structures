// Linked list by Aju

//Delete a node based on the index/position

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

// to Delete - Delete a value in between the nodes
func (l *linkedList) delete(index int) {
	// if the index is more than the list length
	if index >= l.length {
		fmt.Println("Index Out of Bounds")
		return
	}

	//if there is only one node, then set head and tail equal to nil
	if l.length == 1 {
		l.head = nil
		l.tail = nil
		l.length = 0
		return
	}

	//if index is 0, then set the head to the next node
	if index == 0 {
		l.head = l.head.next
		l.length--
		return
	}

	//otherwise
	current := l.head
	for i := 0; i < l.length; i++ {
		if i == index-1 {
			current.next = current.next.next

			//check if the index is our last node, then set the tail to current node
			if index == l.length-1 {
				l.tail = current
			}
			break

		}

		current = current.next
	}
	l.length--
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

	list.delete(3)
	fmt.Println("Length is:", list.len())
	list.display()
	list.delete(2)
	fmt.Println("Length is:", list.len())
	list.display()
	list.delete(3)
	fmt.Println("Length is:", list.len())
	list.display()
	list.delete(0)
	fmt.Println("Length is:", list.len())
	list.display()
	list.delete(25)
	fmt.Println("Length is:", list.len())
	list.display()

}
