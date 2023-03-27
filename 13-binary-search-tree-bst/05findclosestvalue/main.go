//https://github.com/pinglu85/algoExpert/blob/main/Easy/find-closest-value-in-bst.md
//Given a Binary Search Tree and a target integer, we are asked to write a function that is going to return the value in the BST that is closest to the target.
// We can assume there is only one closest value.

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

func (n *node) findClosestValue(target int) int {
	currNode := n
	closest := n.data

	for currNode != nil {
		if absDiff(currNode.data, target) < absDiff(closest, target) {
			closest = currNode.data
		}

		if target < currNode.data {
			currNode = currNode.left
		} else if target > currNode.data {
			currNode = currNode.right
		} else {
			break
		}
	}

	return closest
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
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

	fmt.Println("\nclosest value is", tree1.root.findClosestValue(100))

}

func displayPreorder(t *node) {
	if t == nil {
		return
	}
	fmt.Print(t.data, "-")
	displayPreorder(t.left)
	displayPreorder(t.right)

}
