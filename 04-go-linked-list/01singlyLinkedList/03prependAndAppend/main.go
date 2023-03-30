//Add a node at the end & beginning

package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

//to find length
func (l *linkedList) len() int {
	return l.length
}

//to display
func (l *linkedList) display() {
	temp := l.head //inorder to store the address of head, otherwise head address will be lost and if we print again after two steps ,then the initial head address is lost and not able to print the values.. this code written by myself
	for l.head != nil {
		fmt.Printf("%v ->", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
	l.head = temp
}

//to append - means add data at the end of the linked  list
func (l *linkedList) append(n *node) {
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

//to prepend - means add data to the beginning of the linked  list
func (l *linkedList) prepend(n *node) {
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

	node1 := node{data: 20} //creating node here and passing will cause error if we pass a node multiple times.., so its better to pass values...
	node2 := node{data: 30}
	node3 := node{data: 40}
	node4 := node{data: 50}
	node5 := node{data: 60}
	node6 := node{data: 70}
	node7 := node{data: 80}
	node8 := node{data: 100}
	list := linkedList{}
	list.append(&node1)
	list.append(&node2)
	list.append(&node3)
	list.append(&node4)
	list.append(&node5)
	list.append(&node6)
	list.append(&node7)
	list.prepend(&node8)
	fmt.Println("The length is ", list.len())
	list.display()
	//list.append(&node7)
	fmt.Println("The length is ", list.len())
	list.display()
	// list.prepend(&node8)
	// list.append(&node7)
	// fmt.Println("The length is ", list.len())
	// list.display()

}
