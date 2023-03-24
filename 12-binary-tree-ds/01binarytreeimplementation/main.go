package main

type node struct {
	data        int
	left, right *node
}

type binaryTree struct {
	root *node
}

func (t *binaryTree) insert(value int) *binaryTree {
	//creating a new node and store the value
	//n := &node{data: value}

	if t.root != nil {
		t.root.insertn(value)
	} else {
		t.root = &node{data: value, left: nil, right: nil}
	}
	return t

}

func (n *node)insertn(value int)  {
	if n == nil{
		return
	} else if value > 
}


func main() {

	tree1 := &binaryTree{}
	tree1.insert(50)
}

/*
//stebin babu pgm

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
