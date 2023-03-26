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

func main() {

	tree1 := binaryTree{}

	// tree1.insert(65, "left")
	// tree1.insert(50, "right")
	// tree1.insert(70, "left")
	// tree1.insert(60, "right")
	// tree1.insert(40, "left")
	// tree1.insert(71, "left")
	// tree1.insert(25, "right")
	// tree1.insert(49, "left")
	toggle := 0
	for i := 1; i < 10; i++ {
		tree1.insert(50*i, toggle)
		toggle = 1 - toggle
	}

	fmt.Println("\n Display Inorder traverse is:")
	displayInorder(tree1.root)

	fmt.Println("\n Display Preorder traverse is:")
	displayPreorder(tree1.root)
	fmt.Println("\n Display Postorder traverse is:")
	displayPostorder(tree1.root)

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

/*
//stebin babu pgm - but written in another lgoic by creating class,,, no need of this

package main

import "fmt"

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func display(t *treeNode) {
	if t == nil {
		return
	}
	display(t.left)
	fmt.Print(t.data, "-")
	display(t.right)
}
func (n *treeNode) validate() bool {
	if n.left != nil {
		if n.left.data < n.data {
			return n.left.validate()
		} else {
			return false
		}
	}
	if n.right != nil {
		if n.right.data > n.data {
			return n.right.validate()
		} else {
			return false
		}
	}
	return true
}
func (t *treeNode) getTreeNodeCount() int {
	if t == nil {
		return 0
	}
	return t.left.getTreeNodeCount() + t.right.getTreeNodeCount() + 1
}
func (t *treeNode) TreeDegree() int {
	degree := 0
	if t == nil {
		return degree
	}
	if t.left.TreeDegree() > t.right.TreeDegree() {
		degree = t.left.TreeDegree()
	} else {
		degree = t.right.TreeDegree()
	}
	return degree + 1
}
func main() {
	root := &treeNode{data: 10}
	root.left = &treeNode{data: 9}
	root.right = &treeNode{data: 15}
	root.left.left = &treeNode{data: 20}
	root.left.right = &treeNode{data: 25}
	root.right.left = &treeNode{data: 12}
	root.right.right = &treeNode{data: 30}
	root.left.left.right = &treeNode{data: 50}
	display(root)
	fmt.Println("")
	count := root.getTreeNodeCount()
	fmt.Println("Number of nodes:", count)
	depth := root.TreeDegree()
	fmt.Println("Depth of tree:", depth)
	fmt.Println(root.validate())
}
*/
