package main

import "fmt"

type stack struct {
	top, bottom *node
	length      int
}
type node struct {
	data string
	next *node
}

func (s *stack) push(value string) {
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

}

func (s *stack) pop() {
	fmt.Println(s.top.data)
	if s.length <= 0 {
		fmt.Println("stack is empty")
		return
	}
	if s.length == 1 {
		//fmt.Println(s.top.data)
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

func main() {
	var choice int
	st := stack{}

again:
	fmt.Println("Enter the choice 1 - Push, 2: Pop, 3: Peek, 4:Display, 6:exit")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var data string
		fmt.Println("Enter the value")
		fmt.Scan(&data)
		st.push(data)
		goto again

	case 2:
		st.pop()
		goto again
	case 3:
		st.peek()
		goto again
	}
}
