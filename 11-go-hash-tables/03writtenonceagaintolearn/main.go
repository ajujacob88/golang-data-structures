// This is once again written by myself .. code is same as the 01usinglinkedlistandarray , but here written myself(but same code only) and here no comments are written

package main

import "fmt"

const ArraySize = 8

type hashTable struct {
	array [ArraySize]*bucket
}

type bucketNode struct {
	key  string
	next *bucketNode
}

type bucket struct {
	head *bucketNode
}

func initialise() *hashTable {
	h := &hashTable{}

	for i := range h.array {
		h.array[i] = &bucket{}
	}
	return h
}

func hash(key string) int {
	sum := 0
	for i := range key {
		sum = sum + int(key[i])
	}
	return sum % ArraySize
}

func (h *hashTable) insert(key string) {
	index := hash(key)

	h.array[index].insertb(key)
}

func (b *bucket) insertb(k string) {

	if !b.searchb(k) {
		n := &bucketNode{key: k}
		n.next = b.head
		b.head = n
		//fmt.Println("Added")
	} else {
		fmt.Println("Same key exists")
	}

}

func (b *bucket) searchb(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			//fmt.Println("Same key exists")
			return true

		}
		currentNode = currentNode.next
	}
	return false
}

func (h *hashTable) search(key string) bool {
	index := hash(key)
	return h.array[index].searchb(key)
}

func (h *hashTable) delete(key string) {
	index := hash(key)
	h.array[index].deleteb(key)

}

func (b *bucket) deleteb(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}

}

func main() {
	// hashTest1 := hashTable{}
	// fmt.Println(hashTest1)

	// initialise()
	// fmt.Println(initialise())

	hashTable1 := initialise()

	//testBucket := bucket{}
	//testBucket.insertb("AJU")
	// hashTable1.insert("Aju")
	// hashTable1.insert("Aju")
	// fmt.Println(hashTable1.search("Aju"))

	list := []string{"Aju", "Rahul", "Eldho", "sreejith"}

	for _, v := range list {
		hashTable1.insert(v)
	}

	fmt.Println(hashTable1.search("Anu"))
	fmt.Println(hashTable1.search("Rahul"))
	hashTable1.delete("Rahul")

	fmt.Println(hashTable1.search("Rahul"))

}
