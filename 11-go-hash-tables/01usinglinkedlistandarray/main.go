// https://www.youtube.com/watch?v=zLnJcAt1aKs

// do debug console to find the output.. here we are not printing the output, just analyse the debug console by creating break point. watch the above junmin lee video for understand that..

package main

import (
	"fmt"
)

// ArraySize is the size of the hash table array.
const ArraySize = 7 //we use array size here and there like in hash functions etc, so give it as  a constant and global

// hash table structure - the hashTable will hold an array
type hashTable struct {
	array [ArraySize]*bucket // the hash table(array) will hold the address of the bucket
}

// bucket structure - bucket is a linked list in each of the index of hash table
type bucket struct { //bucket is the linked list
	//length     int
	head *bucketNode
}

// bucket nod is a linked list node that hols the key and next address
type bucketNode struct { // this is the node in the linked list
	key  string //here we call data as key, because it may contains several datas
	next *bucketNode
}

// insert for hashTable structure - insert will take in a key and add it to the hash table array
func (h *hashTable) insert(key string) { // a method of hash table
	index := hash(key) //calling the hash function

	h.array[index].insertb(key) //this insert woulb be a method of a bucket
}

// search for hashTable structure - search will take in a key and return true if that key is stored in the hash table
func (h *hashTable) search(key string) bool {
	index := hash(key)
	return h.array[index].searchb(key)
}

// delete for hashTable structure - delete will take in a key and delete it from the hash table
func (h *hashTable) delete(key string) {
	index := hash(key)
	h.array[index].deleteb(key)
}

// insert for bucket structure - insert will take in a key, create a node with the key and insert the node in the bucket.
func (b *bucket) insertb(k string) {

	if !b.searchb(k) { //meaning b.searchb(k) == false
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("ALready Exists")
	}
}

// search for bucket structure - search will take in a key and return true if the bucket has that key.
func (b *bucket) searchb(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete for bucket structure -  delete will take in a key and delete the node from the bucket
func (b *bucket) deleteb(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
			return
		}

		previousNode = previousNode.next
	}
}

// hash function - get the asci code for each characters and sum it up and divide it by the array size and return the reminder value
func hash(key string) int {
	sum := 0
	for _, v := range key { //to loop through each characters of the key
		sum = sum + int(v)
	}
	return sum % ArraySize
}

// initialise hash table. init will create a bucket in each slot of the hash table.
func initialise() *hashTable {
	result := &hashTable{}
	for i := range result.array { //we need to create bucket 7 times sinse size is defined as 7
		result.array[i] = &bucket{}
	}
	return result
}

func main() {

	//these are for checking while writing the functions and all
	// testHashTable := &hashTable{}
	// fmt.Println(testHashTable)

	// testHashTable1 := initialise()
	// fmt.Println(testHashTable1)

	// fmt.Println(hash("RANDY"))
	// testBucket := &bucket{}
	// testBucket.insertb("RANDY")
	// fmt.Println(testBucket.searchb("RANDY"))
	// fmt.Println(testBucket.searchb("AJU"))

	// testBucket.insertb("RANDY")

	// testBucket.insertb("ANU")
	// fmt.Println(testBucket.searchb("ANU"))
	// testBucket.deleteb("ANU")
	// fmt.Println(testBucket.searchb("ANU"))

	//main function
	hashTable := initialise()
	fmt.Println(hashTable)
	list := []string{"AJU", "ANU", "ELDHO", "RAHUL", "SREEJITH", "NIKHILA", "FEMI", "DEEPA", "MEGHA"}
	for _, v := range list {
		hashTable.insert(v)
	}
	//create a breakpoint at line 137 and use debug console to view the list is stored in hastable or not.... that way we can check the output

	fmt.Println(hashTable.search("ANU"))
	hashTable.delete("ANU") // create a breakpoint at line 143 and check view the deletion process...
	fmt.Println(hashTable.search("ANU"))

}
