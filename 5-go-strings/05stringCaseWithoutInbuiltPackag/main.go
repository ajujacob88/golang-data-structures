// Golang Program to convert a string into Uppercase without using in built function
//https://www.tutorialspoint.com/golang-program-to-convert-a-string-into-uppercase

package main

import (
	"fmt"
	"strings"
)

// function to convert characters to upper case
func ToUpper(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = c - ('a' - 'A')
		}
	}
	return string(b)
}

func main() {
	var input string = "convert this string to uppercase"
	fmt.Println("The given string is:\n", input)
	var output string = ToUpper(input)
	fmt.Println()
	fmt.Println("String obtained by converting the above string to uppercase is:\n", output)

	//if using inbuilt function then
	str := "ajujacob"
	result := strings.ToUpper(str)
	fmt.Println("\nString in uppercase : ", result)
}
