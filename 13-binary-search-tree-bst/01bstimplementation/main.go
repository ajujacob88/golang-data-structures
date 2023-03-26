//https://blog.devgenius.io/golang-data-structures-binary-search-tree-3-5a342d9df2cf

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
		fmt.Println("VALUE NOTT FOUND")
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
		//the logic below is basically the same as above, the only difference is we change the right child of our parent
		//fmt.Println("right.data is:", curr.right.data)
		parent := curr
		curr = curr.right
		rightNode := curr.right

		//fmt.Println("right node is: ", rightNode)

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
			rightNode.left, rightNode.right = nil, nil

		}

	}

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

	tree1.remove(70)

	//display1(os.Stdout, tree1.root, 0, 'M')

	// fmt.Println("\n Display Inorder traverse is:")
	// displayInorder(tree1.root)
	fmt.Println("\n Display Preorder traverse is:")
	displayPreorder(tree1.root)
	// fmt.Println("\n Display Postorder traverse is:")
	// displayPostorder(tree1.root)

	// fmt.Println("")

	// tree1.lookup(50)

	// tree1.lookup(500)

}

// func displayInorder(t *node) {
// 	if t == nil {
// 		return
// 	}
// 	displayInorder(t.left)
// 	fmt.Print(t.data, "-")
// 	displayInorder(t.right)

// }

func displayPreorder(t *node) {
	if t == nil {
		//fmt.Println(t)
		return
	}
	fmt.Print(t.data, "-")
	displayPreorder(t.left)
	displayPreorder(t.right)
	//time.Sleep(1000000000)

}

// func displayPostorder(t *node) {
// 	if t == nil {
// 		return
// 	}
// 	displayPostorder(t.left)
// 	displayPostorder(t.right)
// 	fmt.Print(t.data, "-")
// }

/*
//nor need , other method to print
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

// original before editing
/*
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
		} else {
			//check if the left child is equal to the value.
			//doing this to save the parent node of the value
			if curr.right != nil && curr.right.data == value {
				break
			}
			curr = curr.right
		}
	}
	fmt.Println(curr)
	//if the value we are looking for not found in our tree
	if curr == nil {
		fmt.Println("VALUE NOTT FOUND")
		return
	}

	//check from the parent node is the value is on the left or the right child?
	if curr.left.data == value {
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
			//if the right child has left child, then traverse to the left until we meet null
			for rightNode.left != nil {
				rightNode = rightNode.left
			}
			//change the left node from our parent to the node that we found
			parent.left = rightNode
			//change the left child and the right child to the current child
			rightNode.left, rightNode.right = curr.left, curr.right
			//make sure delet the connection in the current node so it can get pick up by the garbage collection
			curr.left, curr.right = nil, nil
		}
	} else if curr.right.data == value {
		//the logic below is basically the same as above, the only difference is we change the right child of our parent

		parent := curr
		curr = curr.right
		rightNode := curr.right

		if rightNode == nil {
			parent.right = curr.left
			curr.left = nil
			return
		}

		if rightNode.left == nil {
			parent.left = rightNode
			rightNode.left = curr.left
			curr.right, curr.left = nil, nil
		} else {
			for rightNode.left != nil {
				rightNode = rightNode.left
			}
			parent.right = rightNode
			rightNode.left, rightNode.right = curr.left, curr.right

		}

	}

}

*/
