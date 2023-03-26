//https://blog.devgenius.io/golang-data-structures-binary-search-tree-3-5a342d9df2cf
// https://github.com/puneeth8994/binary-tree-go-impl/blob/master/main.go

//this is the same code as 01bstimplementation, but written only largest smallest and insert function is here

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

// FindMax finds the max element in the given BST
func (t *binaryTree) maxElement() int { //actually no need of this func, we can call directly, similar to calling findMin
	return t.root.findMax()
}

func (t *node) findMax() int {
	if t.right == nil {
		return t.data
	}
	return t.right.findMax()
}

// FindMin finds the min element in the given BST
func (t *node) findMin() int {
	if t.left == nil {
		return t.data
	}
	return t.left.findMin()
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

	fmt.Println("The maximum element in the BST is:", tree1.maxElement())
	fmt.Println("The minimum element in the BST is:", tree1.root.findMin())

}

func displayPreorder(t *node) {
	if t == nil {
		return
	}
	fmt.Print(t.data, "-")
	displayPreorder(t.left)
	displayPreorder(t.right)

}
