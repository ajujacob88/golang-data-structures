//https://blog.devgenius.io/golang-data-structures-linked-list-1-2e85a90b2605#:~:text=A%20singly%20linked%20list%20is,which%20is%20the%20next%20node.&text=To%20create%20a%20singly%20linked,we%20will%20store%20our%20value.

package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}
type linkedList struct {
	length int
	head   *node
	tail   *node
}

func (l *linkedList) len() int {
	return l.length
}

func (l *linkedList) display() {
	temp := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%v -> ", temp.data)
		temp = temp.next
	}
	fmt.Println()

}

func (l *linkedList) append(value int) { // append is same as of singly linked list
	n := &node{data: value}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
		l.length++
	}

}

func (l *linkedList) prepend(value int) {
	n := &node{data: value}
	current := l.head
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {

		l.head = n
		l.head.next = current
		l.head.prev = n
		l.length++

	}
}

// func (l *linkedList) insert(index int, value int) {
// 	n := &node{data: value}
// 	if index == 0 {
// 		l.append(value)
// 		return
// 	}
// 	if index == l.length-1 {
// 		l.prepend(value)
// 		return
// 	}

// }

func main() {
	list := linkedList{}
	list.append(25)
	list.append(255)
	list.append(525)
	fmt.Println("Length is:", list.len())
	list.display()
	list.append(5250)
	fmt.Println("Length is:", list.len())
	list.display()
	list.prepend(6666)
	fmt.Println("Length is:", list.len())
	list.display()
	// list.insert(2, 5555)
	// fmt.Println("Length is:", list.len())
	// list.display()
}
