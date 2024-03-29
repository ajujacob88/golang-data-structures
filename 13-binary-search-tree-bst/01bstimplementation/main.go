//https://blog.devgenius.io/golang-data-structures-binary-search-tree-3-5a342d9df2cf
//with referance to the above document, the program written by myself,, there are lot of errors are there in the reference document and many conditions are not checked there.
//But I corrected all those flaws by myself and this code is now fully correct functional....

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
	//create a new node and store the value
	n := &node{data: value}

	//check if we have any nodes in our tree
	if t.root == nil {
		t.root = n
		return
	}

	curr := t.root

	//traverse through the tree
	for curr != nil {
		//if the value is less than the current node, traverse to the left child, else traverse to the right child.
		if n.data < curr.data {
			//if the left node points to nil, insert the new node as the left child
			if curr.left == nil {
				curr.left = n
				break
			}
			curr = curr.left
		} else {
			//if the right node points to nil, insert the new node as the right child
			if curr.right == nil {
				curr.right = n
				break
			}
			curr = curr.right
		}

	}

}

// this lookup is  written by myself with some modifications from lookup function after this, and this code will works fine if no root is present also..
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

/* this is the code i adapted, from this I created the previous lookup function
func (t *binaryTree) lookup(value int) {
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
*/

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
		//if the value is less than current, then go to the left child, else go to the right.
		if value < curr.data {
			//check if the left child is equal to the value
			//doing this to save the parent node of the value
			if curr.left != nil && curr.left.data == value {
				break
			}
			curr = curr.left
		} else if value == curr.data {
			//curr.left = t.root
			break
		} else {
			//check if the left child is equal to the value.
			//doing this to save the parent node of the value
			if curr.right != nil && curr.right.data == value {
				break
			}
			curr = curr.right
		}
	}
	//fmt.Println(curr)
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
			//fmt.Println("in loop 2 - worked fine")
			t.root = rightNode
			t.root.right = rightNode.right
			t.root.left = curr.left

			curr.right, curr.left = nil, nil
		} else {
			parentRightNode := rightNode
			//fmt.Println("in loop 3 - worked fine")
			var temp *node
			for rightNode.left != nil {
				parentRightNode = rightNode
				rightNode = rightNode.left
				//fmt.Println("loop exit rightnode is", rightNode, "rightnode.left is", rightNode.left, "and rightnode.right is", rightNode.right)
				//rightnode is mutable since its * type, so assign the right to a temp variable
				temp = rightNode.right
			}
			// fmt.Println(rightNode)
			//fmt.Println("out loop rightnode is", rightNode, "rightnode.left is", rightNode.left, "and rightnode.right is", rightNode.right)

			t.root = rightNode
			// fmt.Println(t.root)
			// fmt.Println(curr.left)
			// fmt.Println(curr.right)

			t.root.left = curr.left
			//fmt.Println("out loop again is", rightNode, "rightnode.left is", rightNode.left, "and rightnode.right is", rightNode.right)

			t.root.right = curr.right

			//parentRightNode.left = rightNode.right
			parentRightNode.left = temp
			//fmt.Println("rightnode is", rightNode, "rightnode.left is", rightNode.left, "and rightnode.right is", rightNode.right)

			//fmt.Println("t1 root is", t.root)

			curr.left, curr.right = nil, nil
		}

		//check from the parent node is the value is on the left or the right child?
	} else if curr.left != nil && curr.left.data == value {
		//store the parent
		parent := curr
		//store the node we're looking for
		curr = curr.left
		//store the right child from the current node
		rightNode := curr.right

		//if the right child is nil, then move the left child to our current node
		if rightNode == nil {
			parent.left = curr.left
			curr.left = nil
			return
		}

		//if the right child doesnt have any left child, then move the right child to the current node
		if rightNode.left == nil {
			parent.left = rightNode
			rightNode.left = curr.left
			curr.right, curr.left = nil, nil
		} else {
			parentRightNode := rightNode
			//if the right child has left child, then traverse to the left until we meet null
			var temp *node
			for rightNode.left != nil {
				parentRightNode = rightNode
				rightNode = rightNode.left
				temp = rightNode.right
			}
			//change the left node from our parent to the node that we found
			parent.left = rightNode
			//change the left child and the right child to the current child
			rightNode.left, rightNode.right = curr.left, curr.right
			parentRightNode.left = temp
			//make sure delet the connection in the current node so it can get pick up by the garbage collection
			curr.left, curr.right = nil, nil
		}
	} else if curr.right.data == value {
		//store the parent
		parent := curr
		//store the node we're looking for
		curr = curr.right
		//store the right child from the current node
		rightNode := curr.right

		//if the right child is nil, then move the left child to our current node
		if rightNode == nil {
			parent.right = curr.left
			curr.left = nil
			return
		}

		//if the right child doesnt have any left child, then move the right child to the current node
		if rightNode.left == nil {
			parent.right = rightNode
			rightNode.left = curr.left
			curr.right, curr.left = nil, nil
		} else {
			parentRightNode := rightNode
			//if the right child has left child, then traverse to the left until we meet null
			var temp *node
			for rightNode.left != nil {
				parentRightNode = rightNode
				rightNode = rightNode.left
				temp = rightNode.right
			}
			//change the left node from our parent to the node that we found
			parent.right = rightNode
			//change the left child and the right child to the current child
			rightNode.left, rightNode.right = curr.left, curr.right
			parentRightNode.left = temp
			//make sure delet the connection in the current node so it can get pick up by the garbage collection
			curr.left, curr.right = nil, nil
		}

	}

}

//to find the max and min
// https://github.com/puneeth8994/binary-tree-go-impl/blob/master/main.go

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

	tree1.remove(65)

	fmt.Println("\n Display Inorder traverse is:")
	displayInorder(tree1.root)
	fmt.Println("\n Display Preorder traverse is:")
	displayPreorder(tree1.root)
	fmt.Println("\n Display Postorder traverse is:")
	displayPostorder(tree1.root)

	fmt.Println("")

	tree1.lookup(93)

	tree1.lookup(500)
	//BST status true means the tree is a BST and False means the tree is Not a BST..
	fmt.Println("BST Status:  ", tree1.ValidateBst())

	fmt.Println("The maximum element in the BST is:", tree1.maxElement())
	fmt.Println("The minimum element in the BST is:", tree1.root.findMin())
	//display1(os.Stdout, tree1.root, 0, 'M')

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

/*
// nor need , other method to print
func display1(w io.Writer, node *node, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	display1(w, node.left, ns+2, 'L')
	display1(w, node.right, ns+2, 'R')
}
*/
