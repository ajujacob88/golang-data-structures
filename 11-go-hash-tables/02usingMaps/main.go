package main

import "fmt"

func main() {
	fmt.Println("Hi Hello")
	hashTable()
	fmt.Println("Stored key and values\n")
	fmt.Println(newHash, "\n")
	fmt.Println("New hash table with index\n")
	fmt.Println(newTable)
}

type Hash struct {
	key   []int
	value []string
}
type Node struct {
	data string
	next *Node
}
type linkedlist struct {
	head *Node
	tail *Node
}

var list = linkedlist{}
var newHash = Hash{}
var newTable = [10]string{}
var i int = 0

func hashTable() {
	var limit, key int
	var value string
	fmt.Println("Enter the limit")
	fmt.Scan(&limit)
	for i := 0; i < limit; i++ {
		fmt.Printf("Enter the intiger key : ")
		fmt.Scan(&key)
		fmt.Printf("Enter the value for key %v :", key)
		fmt.Scan(&value)
		newHash.insert(key, value)
		newHash.hashing(key, value)
	}
}

func (h *Hash) insert(key int, value string) {
	h.key = append(h.key, key)
	h.value = append(h.value, value)
}

func (h *Hash) hashing(key int, value string) {

	index := key % len(newTable)
	if newTable[index] == "" {
		newTable[index] = value
	} else {
		var option int
		fmt.Println("Sorry the index is already filled\nChoose>>>\n1.Chaining\n2.Linear Probing\n3.Quadratic Probing\n4.Double hashing")
		fmt.Scan(&option)
		switch option {
		case 1:
			list.insert(value)
		case 2:
			linear(value, key)
		case 3:
			quadratic(value, index)
		case 4:
			doubleHashing(value, index, i, key)
		}
	}
}

// Collision Handling

// Chaining >>> creating linked list

func (l *linkedlist) insert(value string) {
	node := &Node{data: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	return

}

// Linear Probing

func linear(value string, key int) {
	index := (key + 1) % len(newTable)
	if newTable[index] == "" {
		newTable[index] = value
	} else {
		linear(value, key)
	}
}

// Quadratic Probing

func quadratic(value string, index1 int) {
	index := (index1 + (index1 * index1)) % len(newTable)
	if newTable[index] == "" {
		newTable[index] = value
	} else {
		quadratic(value, index)
	}
}

// Double Hashing

func doubleHashing(value string, index1, i, key int) {
	index2 := 7 - ((key) % 7)
	index := (index1 + i*(index2))
	i++
	if newTable[index] == "" {
		newTable[index] = value
	} else {
		doubleHashing(value, index, i, key)
	}
}
