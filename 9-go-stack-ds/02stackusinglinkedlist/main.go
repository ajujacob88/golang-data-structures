// https://blog.devgenius.io/golang-data-structures-stack-and-queue-2-5a6822baeb0d

package main

import "fmt"

type node struct {
	next *node
	data string
}

type stack struct { //this is the linked list
	top, bottom *node
	length      int
}

func main() {
	var choice int
	st := stack{}

again:
	fmt.Println("Enter choice: 1 - Push, 2: Pop, 3: Peek, 4: Display")
	fmt.Scanln(&choice)

	for choice != 6 {
		switch choice {
		case 1:
			var data string
			fmt.Println("Enter the data")
			fmt.Scanln(&data)
			st.push(data)
			goto again
		case 2:
			st.pop()
			goto again
		case 3:
			st.peek()
			goto again

		case 4:
			st.display()
			goto again
		default:
			fmt.Println("Invalid choice entered")
			goto again
		}
	}

}

func (s *stack) push(value string) {
	n := &node{data: value}

	if s.length == 0 {
		s.bottom = n
	} else {
		n.next = s.top
	}
	s.top = n
	s.length++

	//or below done myself
	/*
		n := &node{data: value}
		if s.length == 0 {
			s.top = n
			s.bottom = n

			s.length++
		} else {
			n.next = s.top
			s.top = n
			s.length++
		}
	*/
}

func (s *stack) pop() {
	if s.length <= 0 {
		fmt.Println("The stack is empty")
		return
	}
	if s.length == 1 {
		s.top = nil
		s.bottom = nil
	} else {
		top := s.top
		s.top = top.next
		top.next = nil
	}
	s.length--

}

func (s *stack) peek() {
	fmt.Println(s.top.data)
}

func (s *stack) display() {

	currentNode := s.top

	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}

}

/*
Applications of Stack
The Stack has a lot of applications as follows

To implement undo feature
Build compilers
Evaluate Expressions like Balance brackets
Reverse something
Navigation (web browsers)

*/
