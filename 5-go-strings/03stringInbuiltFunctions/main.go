// https://pkg.go.dev/strings#Compare
package main

import (
	"fmt"
	"strings"
)

func main() {
	//Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b. func Compare(a, b string) int
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	//Contains reports whether substr is within s.  func Contains(s, substr string) bool
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	//ContainsAny reports whether any Unicode code points in chars are within s. func ContainsAny(s, chars string) bool
	fmt.Println("contains Any")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))

	//Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s. func Count(s, substr string) int
	fmt.Println("Count")
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("malayalam", "y"))
	fmt.Println(strings.Count("malayalam", "a"))
	fmt.Println(strings.Count("five", "")) // before & after each rune

	//Clone returns a fresh copy of s. func Clone(s string) string

	//Concatenating Strings - The strings package includes a method join for concatenating multiple strings − strings.Join(sample, " ")
	var str1 string = "aju"
	var str2 string = "jacob"
	fmt.Println(str1)
	fmt.Println(str2)
	str3 := str1 + str2
	fmt.Println(str3)

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, " nikhil "))

	//to upper case func ToUpper(s string) string
	var str string = "Hello World"

	result := strings.ToUpper(str)
	fmt.Println("String in uppercase : ", result)

	//to lower case - func ToLower(s string) string
	fmt.Println(strings.ToLower("GoPher"))
}
