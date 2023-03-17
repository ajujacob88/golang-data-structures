// In GO maps, the key is map will be automatically sorted based on the key values during runtime,,, so if we want to sort based on values, then we can use this bubble sort

package main

import "fmt"

func main() {
	map1 := map[string]int{"aju": 33, "rahul": 32, "eldho": 31, "nikki": 30, "sree": 29}
	fmt.Println(map1)

	fmt.Println(BubblesortByValue(map1))

}

func BubblesortByValue(data map[string]int) map[string]int {
	for _, val := range data {

	}
}
