//mutable and immutable

package main

import (
	"fmt"
)

func main() {
	var str string = "aju jacobb"
	var ptr = &str
	fmt.Println(str)
	fmt.Println(ptr)
	//str = strings.Replace(str, "j", "n", 1)
	str = "anu"
	ptr1 := &str
	fmt.Println(str)
	fmt.Println(ptr1)

	//string is immutable data type , means we cant change individual chars in a string ,eg we cant assign a value to test1[0] or test1[1] etc if its a string

	test1 := "aju"
	test2 := test1
	test2 = "hi"
	fmt.Println(test1)
	fmt.Println(test2)
	//test1[0] = "s" we cant do like this because string is immutable

	//we can change values of str5 from str6 if we use pointers
	str5 := "rahul"
	str6 := &str5
	*str6 = "jacob"
	fmt.Println(str5)
	fmt.Println(*str6) //printing the value
	fmt.Println(&str5)
	fmt.Println(str6)

	//variable is immutable data type
	test3 := 25
	test4 := test3
	test4 = 50
	fmt.Println(test3)
	fmt.Println(test4)

	//slice is MUTABLE DATATYPE
	test5 := []int{20, 30, 40}
	test6 := test5 //Here test 5 and test 6 is pointing to the same variable..,if we change any one means we are changing both of them, because its just pointing.
	test6[0] = 5500
	fmt.Println(test5)
	fmt.Println(test6)

	// MAP IS ALSO MUTABLE DATATYPE
	var map1 map[string]int = map[string]int{"hello": 2, "hi": 3}
	map2 := map1
	map2["aju"] = 35
	fmt.Println(map1)
	fmt.Println(map2)

	// ARRAY IS ALSO MUTABLE DATATYPE, BUT HERE THE array will get a copy of array1, here its not just pointing, but it will get a copy, but anyway array is also mutable...
	array1 := [5]int{20, 30, 40}
	array2 := array1
	array2[0] = 6000
	fmt.Println(array2)
	fmt.Println(array1)

	// mutable functions, if we call a slice into a function also, then also same things happens,, here also recslice pointing to the copy only

	slice1 := []int{2, 3, 4}
	fmt.Println("orig slice: ", slice1)
	changeslice(slice1)
	fmt.Println("modified via function: ", slice1)
	var pointr *[]int
	pointr = &slice1
	changesliceviapointer(pointr)
	fmt.Println(slice1)

}

func changeslice(recslice []int) {
	recslice[0] = 500
}

func changesliceviapointer(modifyslice *[]int) {
	fmt.Println(modifyslice)
	(*modifyslice)[1] = 255
}
