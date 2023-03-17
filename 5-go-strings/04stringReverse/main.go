package main

import "fmt"

func main() {

	reversed := reverseString("hello")
	fmt.Println(reversed)

}

func reverseString(s string) string {
	var reversed string

	for i := len(s) - 1; i >= 0; i-- {
		reversed = reversed + string(s[i]) - (32)

		// fmt.Println(s[i])
		// fmt.Println(string(s[i]))
	}

	return reversed
}

// Another method - more standard method i think, using rune..but i didnt learned it till now
// Golang program to reverse a string - https://www.geeksforgeeks.org/how-to-reverse-a-string-in-golang/
/*
package main

// importing fmt
import "fmt"

// function, which takes a string as
// argument and return the reverse of string.
func reverse(s string) string {
    rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }

    // return the reversed string.
    return string(rns)
}

func main() {

    // Reversing the string.
    str := "Geeks"

    // returns the reversed string.
    strRev := reverse(str)
    fmt.Println(str)
    fmt.Println(strRev)
}
*/

/*

package main

import "fmt"

func main() {

	var str string = "aju"

	fmt.Println(str)

	str = "jacob"
	fmt.Println(str)
	j := len(str)

	var str1 [5]string
	str1[1] = "a"

	//str = string(str[3]) + string(str[4])

	for i := 0; i < j; i++ {

		str = string(str[j-1]) + string(str[4])
	}
	fmt.Println(str)
	//j := len(str)
	// for i := j - 1; i >= 0; i-- {

	// 	fmt.Print(string(str[i]))

	// 	//fmt.Println(j)
	// 	//string(str[i]) = "a"
	// }
	// 	var temp string
	// 		temp[0] = "a"
	// 	// for i := 0; i < j; i++ {
	// 	// 	str1 := str[i]
	// 	// 	fmt.Println(string(str1))
	// 	// 	str[i] = str1[0]
	// 	// }
	// 	for i := 0; i <= j; i++ {
	// 		temp = string(str[j-1])
	// 		fmt.Println(temp)
	// 		j--

	// 	}

}
*/
