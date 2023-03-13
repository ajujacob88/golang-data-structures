package main

import "fmt"

func linearSearch(inp string, key string) int {
	for i := 0; i < len(inp); i++ {
		if string(inp[i]) == key {
			return i
		}
	}
	return -1
}

func main() {
	str := "aju jacob"
	out := linearSearch(str, "c")
	fmt.Println(out)

	index := FindIndex("aju jacob", "c")

	fmt.Println(index)
}

//other method

func FindIndex(input string, character string) int {

	for i, v := range input {
		if string(v) == character {
			return i
		}
	}
	return -1
}
