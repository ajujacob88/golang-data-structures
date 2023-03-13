package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter a string")
	fmt.Scan(&str)

	var flag int
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			flag = 1
			break
		}
	}

	if flag == 1 {
		fmt.Println("Entered string is Not a Palindrome")
	} else {
		fmt.Println("Entered string is Palindrome")
	}

}
