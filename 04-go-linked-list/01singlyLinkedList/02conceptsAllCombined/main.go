// https://hackthedeveloper.com/golang-linked-list/
// https://www.interviewbit.com/blog/types-of-linked-list/#:~:text=list%20using%20indexes.-,Head%20and%20Tail%20pointer,the%20list%20in%20constant%20time.

package main

import "fmt"

// Node Struct
type node struct {
	data int
	next *node
}

// Linked List Struct
type linkedList struct {
	length int
	head   *node
	tail   *node
}

//Length Of The Linked List.. This method returns the Length of the Linked List.
func (l linkedList) Len() int {
	return l.length
}

//Linked List Display
func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v ->", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

//Linked List PushBack.. The PushBack Method takes a node as input and links it to the linked list.
func (l *linkedList) PushBack(n *node) { //this is the append,, adding node at the end
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

//this is copied from medium
func (l *linkedList) prepend(n *node) { //adding node at the beginning
	//create new node
	// el := &node{
	// 	data: value,
	// }
	// Check if there are nodes in our current list, if no then add new node to the tail
	// else, set the next to current head and set head to new node
	if l.length == 0 {
		l.tail = n
	} else {
		n.next = l.head
	}
	l.head = n
	l.length += 1
}

func (l *linkedList) prepend2(value int) {
	//create new node
	el := &node{data: value}
	// Check if there are nodes in our current list, if no then add new node to the tail
	// else, set the next to current head and set head to new node
	if l.length == 0 {
		l.tail = el
	} else {
		el.next = l.head
	}
	l.head = el
	l.length += 1
}

func main() {
	node1 := &node{data: 20}
	list := linkedList{} //Linked List Struct Initialization In Main Function.
	list.PushBack(node1) //Main Function Code For The Pushback Function.

	node2 := &node{data: 320}
	node3 := &node{data: 400}
	node4 := &node{data: 50}
	node5 := &node{data: 600}
	node6 := &node{data: 1220}
	list.PushBack(node2)
	list.PushBack(node3)
	list.PushBack(node4)
	list.PushBack(node5)
	list.prepend(node6)

	list.prepend2(2500)
	list.prepend2(2500)
	fmt.Println("Length = ", list.Len())
	list.Display()

}
