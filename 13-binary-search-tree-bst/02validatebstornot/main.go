package main

import (
	"fmt"
	"math"
)

type node struct {
	data        int
	left, right *node
}

type binaryTree struct {
	root *node
}

func (t *binaryTree) insert(value int, toggle int) {
	n := &node{data: value}

	if t.root == nil {
		t.root = n
		return
	}

	curr := t.root

	//traverse through the tree
	for curr != nil {
		if toggle == 0 {
			if curr.left == nil {
				curr.left = n
				break
			} else if curr.right == nil {
				curr.right = n
				break
			}
			curr = curr.left
		} else if toggle == 1 {
			if curr.right == nil {
				curr.right = n
				break
			} else if curr.left == nil {
				curr.left = n
				break
			}
			curr = curr.right
		}

	}

}

//validation reference: https://levelup.gitconnected.com/a-crash-course-on-binary-search-trees-in-golang-41251020477d

func (tree *binaryTree) ValidateBst() bool {
	return tree.root.validateBstNode(math.MinInt32, math.MaxInt32)
}

func (tree *node) validateBstNode(minValue, maxValue int) bool {

	if tree.data < minValue || tree.data >= maxValue { //the condition is checking here
		return false
	}

	if tree.left != nil && !tree.left.validateBstNode(minValue, tree.data) {
		return false
	}
	if tree.right != nil && !tree.right.validateBstNode(tree.data, maxValue) {
		return false
	}
	return true
}

/*
//same code above just written for learning or understanding the concepts by printing every line..

	func (tree *node) validateBstNode(minValue, maxValue int) bool {
		fmt.Println("\nmin value 00 is", minValue, " max value 00 is", maxValue, "tree is", tree, "tree data is", tree.data)
		if tree.data < minValue || tree.data >= maxValue { //the condition is checking here
			fmt.Println("min value 0 is", minValue, " max value 0 is", maxValue, "tree is", tree, "tree data is", tree.data)
			return false
		}

		if tree.left != nil && tree.left.validateBstNode(minValue, tree.data) == false {
			fmt.Println("min value 1 is", minValue, " max value 1 is", maxValue, "tree is", tree, "tree data is", tree.data)
			return false
		}
		if tree.right != nil && tree.right.validateBstNode(tree.data, maxValue) == false {
			fmt.Println("min value 2 is", minValue, " max value 2 is", maxValue, "tree is", tree, "tree data is", tree.data)
			return false
		}
		fmt.Println("min value 3 is", minValue, " max value 3 is", maxValue, "tree is", tree, "tree data is", tree.data)
		return true
	}
*/
func main() {

	tree1 := binaryTree{}

	// toggle := 0

	// for i := 1; i < 10; i++ {
	// 	tree1.insert(50*i, toggle)
	// 	toggle = 1 - toggle
	// }

	tree1.insert(65, 0)
	tree1.insert(40, 0)
	tree1.insert(70, 0)
	tree1.insert(75, 1)
	tree1.insert(68, 1)
	tree1.insert(25, 0)
	tree1.insert(42, 0)

	// tree1.insert(50, 0)
	// tree1.insert(40, 0)
	// tree1.insert(60, 0)

	fmt.Println("\n Display Preorder traverse is:")
	displayPreorder(tree1.root)
	//BST status true means the tree is a BST and False means the tree is Not a BST..
	fmt.Println("\nBST Status:  ", tree1.ValidateBst())

}

func displayPreorder(t *node) {
	if t == nil {
		return
	}
	fmt.Print(t.data, "-")
	displayPreorder(t.left)
	displayPreorder(t.right)

}
