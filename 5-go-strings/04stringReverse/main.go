package main

import "fmt"

func main() {

	reversed := reverseString("hello")
	fmt.Println(reversed)

}

func reverseString(s string) string {
	var reversed string

	for i := len(s) - 1; i >= 0; i-- {
		reversed = reversed + string(s[i])

		// fmt.Println(s[i])
		// fmt.Println(string(s[i]))
	}

	return reversed
}
