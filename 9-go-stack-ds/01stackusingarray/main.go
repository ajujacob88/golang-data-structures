package main

import "fmt"

type stack struct {
	limit    int
	stackarr []int
	top      int
}

func main() {
	var choice, limit int
	st := stack{}
	fmt.Println("Enter the limit of the stack : ")
	fmt.Scanln(&limit)
	st.stack(limit)
again:
	fmt.Println("Enter choice: 1 - Push, 2: Pop, 3: Peek, 4:Display, 5: is Full or not, 6:exit")
	fmt.Scanln(&choice)

	for choice != 6 {
		switch choice {
		case 1:
			data := 0
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
		case 5:
			fmt.Println(st.isFull())
			goto again
		case 6:

		default:
			fmt.Println("Invalid choice entered")
			goto again
		}
	}
}

func (s *stack) stack(limit int) {
	s.limit = limit
	s.top = -1
}

func (s *stack) push(data int) {

	s.top++
	s.stackarr = append(s.stackarr, data)
	//s.stackarr[s.top] = data
}

func (s *stack) isEmpty() bool {
	// if s.stackarr == nil {
	// 	return true
	// } else {
	// 	return false
	// }
	if s.top < 0 {
		return true
	} else {
		return false
	}
}

func (s *stack) pop() {
	if !s.isEmpty() { // same as s.isEmpty() == false
		fmt.Println(s.stackarr[s.top])
		// s.stackarr[s.top] = 0
		s.top--
	} else {
		fmt.Println("Stack is emptyy")
	}

}

func (s *stack) peek() {
	fmt.Println(s.stackarr[s.top])
}

func (s *stack) display() {
	for i := s.top; i >= 0; i-- {
		fmt.Println(s.stackarr[i])
	}
}

func (s *stack) isFull() bool {
	if s.top == s.limit-1 {
		return true
	} else {
		return false
	}
}
