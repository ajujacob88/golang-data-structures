// https://www.youtube.com/watch?v=nL7BHR5vJDc&list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6&index=8

package main

import "fmt"

//AlphabetSize is the number of possible characters in the trie.
const AlphabetSize = 26

// struct for a Node - node represents each node in the trie
// The node first needs to hold an array and the array is going to be a size of 26, which is the size of the alphabets.
//and each index of that array is going to hold a pointer pointing to the child
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

//struct for a Trie.. The trie is going to hold a pointer pointing to the root node of that child
// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

//initTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//insert method - Insert will take in a word and add it to the trie.
func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a' // this is to change the character to 0 to 26 index,, just pritn this seperately to get an idea or watch 15:20 minutes in Junmin Lee video
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//search method - search will take a word and Return True if that word is included in the trie.
func (t *Trie) search(word string) bool {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a' // this is to change the character to 0 to 26 index,, just pritn this seperately to get an idea or watch 15:20 minutes in Junmin Lee video
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	// testTrie := InitTrie()
	// fmt.Println(testTrie)
	// fmt.Println(testTrie.root)

	myTrie := InitTrie()
	// myTrie.Insert("aju")
	// //myTrie.search("aju")
	// fmt.Println(myTrie.search("aju"))

	sliceToInsert := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}
	for _, v := range sliceToInsert {
		myTrie.Insert(v)
	}
	fmt.Println(myTrie.search("aju"))
	fmt.Println(myTrie.search("oreo"))
	//fmt.Println(myTrie.root)

}
