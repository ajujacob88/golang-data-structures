package main

import "fmt"

func main() {

	mapList2 := []map[string]int{
		{"name": 25, "age": 52},
		{"name": 18, "age": 58},
		{"name": 20, "age": 45},
	}

	for i := 0; i < len(mapList2)-1; i++ {
		for j := 0; j < len(mapList2)-i-1; j++ {
			// if strings.Compare(mapList2[j+1]["name"], mapList2[j]["name"]) >= 0 {
			// 	mapList2[j], mapList2[j+1] = mapList2[j+1], mapList2[j]

			// }

			if mapList2[j+1]["name"] < mapList2[j]["name"] {
				mapList2[j], mapList2[j+1] = mapList2[j+1], mapList2[j]

			}

		}
	}
	fmt.Println("sorted: ", mapList2)

}

/* from chatgpt - read aju
The Go programming language specification does not guarantee any specific order of iteration when ranging over a map.
Therefore, the order of the printed output may vary between different Go compilers or even between different runs of the same code.

However, in some Go compilers, including the official Go compiler and some third-party compilers, the implementation of the map type iterates over the keys in a sorted order by default.
This means that when you print the map, the keys will appear in alphabetical order.

It is important to note that the order of the keys is not guaranteed by the language specification and should not be relied upon in your code.
If you need to iterate over the map in a specific order, you should sort the keys explicitly before iterating over them.


To sort the keys of a map, you can follow these steps:

Extract the keys from the map into a slice.
Sort the slice using the sort package.
Iterate over the sorted slice, accessing the values in the original map using the keys.
Here's an example code snippet that demonstrates how to sort the keys of a map in Go:


import (
    "fmt"
    "sort"
)

func main() {
    m := map[string]int{"foo": 1, "bar": 2, "baz": 3}

    // Step 1: Extract the keys from the map into a slice.
    var keys []string
    for k := range m {
        keys = append(keys, k)
    }

    // Step 2: Sort the slice using the sort package.
    sort.Strings(keys)

    // Step 3: Iterate over the sorted slice, accessing the values in the original map using the keys.
    for _, k := range keys {
        fmt.Printf("key=%s value=%d\n", k, m[k])
    }
}

In this example, the keys slice is created by iterating over the keys of the m map and appending them to the slice.
The sort.Strings function is then used to sort the keys slice in alphabetical order.
Finally, the sorted keys are iterated over using a for loop and the corresponding values are accessed from the original m map.

*/
