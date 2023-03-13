// Go program to illustrate how to compare
// string using compare() function

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Println("Enter the first string:")
	fmt.Scan(&str1)
	fmt.Println("Enter the second string")
	fmt.Scan(&str2)

	if strings.Compare(str1, str2) == 0 {
		fmt.Println("Both the strings are same")
	} else {
		fmt.Println("Both the strings are different")
	}
}
