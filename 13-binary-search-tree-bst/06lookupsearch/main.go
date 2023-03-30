//This was the same code as in 01bstimplementation, just seperated the lookup function in 01bstimplementation

package main

import (
	"fmt"
)

type node struct {
	data        int
	left, right *node
}

type binaryTree struct {
	root *node
}

func (t *binaryTree) insert(value int) {
	n := &node{data: value}

	if t.root == nil {
		t.root = n
		return
	}

	curr := t.root

	for curr != nil {
		if n.data < curr.data {
			if curr.left == nil {
				curr.left = n
				break
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = n
				break
			}
			curr = curr.right
		}

	}

}

// this lookup is  written by myself with some modifications from lookup2 function, and this code will works fine if no root is present also..
func (t *binaryTree) lookup(value int) {
	if t.root == nil {
		fmt.Println("The value is not present")
		return
	}
	curr := t.root
	for curr != nil {
		if value < curr.data {
			if value == curr.data {
				fmt.Println("The value is present")
				return
			}
			curr = curr.left
		} else {
			if value == curr.data {
				fmt.Println("The value is present")
				return
			}
			curr = curr.right
		}
	}
	fmt.Println("The value is Not present")

}

func main() {

	tree1 := binaryTree{}

	tree1.insert(65)
	tree1.insert(50)
	tree1.insert(70)
	tree1.insert(60)
	tree1.insert(40)
	tree1.insert(81)
	tree1.insert(25)
	tree1.insert(49)
	tree1.insert(61)
	tree1.insert(66)

	tree1.insert(58)
	tree1.insert(59)
	tree1.insert(67)
	tree1.insert(73)
	tree1.insert(90)
	tree1.insert(93)
	tree1.insert(70)

	fmt.Println("\n Display Preorder traverse is:")
	displayPreorder(tree1.root)

	fmt.Println("")

	tree1.lookup2(73)
	tree1.lookup(73)

	//lookup2 will also runs fine if no data is in root node
	tree2 := &binaryTree{}
	//tree2.lookup(25)
	tree2.lookup(25)

}

func displayPreorder(t *node) {
	if t == nil {
		return
	}
	fmt.Print(t.data, "-")
	displayPreorder(t.left)
	displayPreorder(t.right)

}

//this lookup2 is copied from other sites, but the previous function lookup is the modified verision of this which runs fine...

func (t *binaryTree) lookup2(value int) {
	//check if we have only one node
	if t.root.left == nil && t.root.right == nil {
		fmt.Println("value found?", t.root.data == value, t.root)
		return
	}

	curr := t.root
	foundStatus := false

	//traverse through the tree
	for curr != nil {
		if curr.data == value {
			foundStatus = true
			break
		}
		if value < curr.data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	// check if the current node value is equal to value we want to search
	if curr != nil && curr.data == value {
		fmt.Println(value, "-", foundStatus, "Left", curr.left, "Right", curr.right)
	} else {
		fmt.Println(value, "-", foundStatus)
	}

}
