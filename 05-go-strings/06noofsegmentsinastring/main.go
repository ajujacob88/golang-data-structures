/*
434. Number of Segments in a String
Easy
655
1.2K
Companies
Given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.



Example 1:

Input: s = "Hello, my name is John"
Output: 5
Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
Example 2:

Input: s = "Hello"
Output: 1

https://leetcode.com/problems/number-of-segments-in-a-string/description/
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var s = "Hello, my name is John"

	var s1 = "Hello,       my      name      is      John"

	var s3 = "       "

	fmt.Println(countSegments(s))
	fmt.Println(countSegments(s1))
	fmt.Println(countSegments(s3))
}

func countSegments(s string) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}
	flag := 0
	str := []byte(s)

	for i := 0; i < len(str); i++ {
		if s[i] == ' ' {
			for s[i] == ' ' {
				i++
			}
			flag++
		}
	}

	return flag + 1

	// flag = strings.Count(s," ")
	//     flag++
	// }
	// if flag == 0{
	//     return 0
	// }

	// x := strings.Compare(s,"")
	// if x == 0{
	//     return 0
	// }

}
