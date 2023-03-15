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
