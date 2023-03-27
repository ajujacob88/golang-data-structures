//https://blog.devgenius.io/golang-data-structures-binary-search-tree-3-5a342d9df2cf
//with referance to the above document, the program written by myself,, there are lot of errors are there in the reference document and many conditions are not checked there.
//But I corrected all those flaws by myself and this code is now fully correct functional....

//All operations combined, this is same as previous programs, but here deleted the comments and combined all operations

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

func (t *binaryTree) lookup(value int) {
	if t.root.left == nil && t.root.right == nil {
		fmt.Println("value found?", t.root.data == value, t.root)
		return
	}

	curr := t.root
	foundStatus := false

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

	if curr != nil && curr.data == value {
		fmt.Println(value, "-", foundStatus, "Left", curr.left, "Right", curr.right)
	} else {
		fmt.Println(value, "-", foundStatus)
	}

}

// node Deletion

func (t *binaryTree) remove(value int) {
	curr := t.root
	//if we dont have any root, then returns
	if curr == nil {
		return
	}
	//if we have only root, then check the root
	if curr.right == nil && curr.left == nil {
		if curr.data == value {
			t.root = nil
		}
		return
	}

	//traverse through the tree
	for curr != nil {
		if value < curr.data {
			if curr.left != nil && curr.left.data == value {
				break
			}
			curr = curr.left
		} else if value == curr.data {
			break
		} else {
			if curr.right != nil && curr.right.data == value {
				break
			}
			curr = curr.right
		}
	}
	//if the value we are looking for not found in our tree
	if curr == nil {
		fmt.Println("VALUE NOT FOUND")
		return
	}

	if curr.data == value { //this is the case when the root node is deleted,, this is written by myself and works fine

		rightNode := curr.right

		if rightNode == nil {
			t.root = curr.left
			curr.left = nil //inorder to pick this up by garbage collector
			return
		}

		if rightNode.left == nil {
			t.root = rightNode
			t.root.right = rightNode.right
			t.root.left = curr.left

			curr.right, curr.left = nil, nil
		} else {
			parentRightNode := rightNode
			var temp *node
			for rightNode.left != nil {
				parentRightNode = rightNode
				rightNode = rightNode.left
				//rightnode is mutable since its * type, so assign the right to a temp variable
				temp = rightNode.right
			}

			t.root = rightNode

			t.root.left = curr.left

			t.root.right = curr.right

			parentRightNode.left = temp

			curr.left, curr.right = nil, nil
		}

	} else if curr.left != nil && curr.left.data == value {
		parent := curr
		curr = curr.left
		rightNode := curr.right

		if rightNode == nil {
			parent.left = curr.left
			curr.left = nil
			return
		}

		if rightNode.left == nil {
			parent.left = rightNode
			rightNode.left = curr.left
			curr.right, curr.left = nil, nil
		} else {
			parentRightNode := rightNode
			var temp *node
			for rightNode.left != nil {
				parentRightNode = rightNode
				rightNode = rightNode.left
				temp = rightNode.right
			}
			parent.left = rightNode
			rightNode.left, rightNode.right = curr.left, curr.right
			parentRightNode.left = temp
			curr.left, curr.right = nil, nil
		}
	} else if curr.right.data == value {
		parent := curr
		curr = curr.right
		rightNode := curr.right

		if rightNode == nil {
			parent.right = curr.left
			curr.left = nil
			return
		}

		if rightNode.left == nil {
			parent.right = rightNode
			rightNode.left = curr.left
			curr.right, curr.left = nil, nil
		} else {
			parentRightNode := rightNode
			var temp *node
			for rightNode.left != nil {
				parentRightNode = rightNode
				rightNode = rightNode.left
				temp = rightNode.right
			}
			parent.right = rightNode
			rightNode.left, rightNode.right = curr.left, curr.right
			parentRightNode.left = temp
			curr.left, curr.right = nil, nil
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
	fmt.Println("\n Display Preorder traverse before deletion:")
	displayPreorder(tree1.root)
	tree1.remove(65)

	fmt.Println("\n Display Inorder traverse is:")
	displayInorder(tree1.root)
	fmt.Println("\n Display Preorder traverse is:")
	displayPreorder(tree1.root)
	fmt.Println("\n Display Postorder traverse is:")
	displayPostorder(tree1.root)

	fmt.Println("")

	tree1.lookup(50)

	tree1.lookup(500)
	//BST status true means the tree is a BST and False means the tree is Not a BST..
	fmt.Println("BST Status:  ", tree1.ValidateBst())

	fmt.Println("The maximum element in the BST is:", tree1.maxElement())
	fmt.Println("The minimum element in the BST is:", tree1.root.findMin())

	fmt.Println("\nclosest value is", tree1.root.findClosestValue(100))

}

func displayInorder(t *node) {
	if t == nil {
		return
	}
	displayInorder(t.left)
	fmt.Print(t.data, "-")
	displayInorder(t.right)

}

func displayPreorder(t *node) {
	if t == nil {
		return
	}
	fmt.Print(t.data, "-")
	displayPreorder(t.left)
	displayPreorder(t.right)

}

func displayPostorder(t *node) {
	if t == nil {
		return
	}
	displayPostorder(t.left)
	displayPostorder(t.right)
	fmt.Print(t.data, "-")
}

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
