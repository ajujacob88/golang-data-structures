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

func main() {

	tree1 := binaryTree{}

	tree1.insert(65)
	tree1.insert(50)
	tree1.insert(70)
	tree1.insert(60)
	tree1.insert(40)
	tree1.insert(71)
	tree1.insert(25)
	tree1.insert(49)

	//display1(os.Stdout, tree1.root, 0, 'M')
	display(tree1.root)

}

func display(t *node) {
	if t == nil {
		return
	}
	display(t.left)
	fmt.Print(t.data, "-")
	display(t.right)
}

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
