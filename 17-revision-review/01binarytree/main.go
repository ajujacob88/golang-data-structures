/*package main

import "fmt"

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
	tree1 := &binaryTree{}
	tree1.insert(25, 0)
	tree1.insert(50, 1)
	tree1.insert(75, 0)

	toggle := 0
	for i := 1; i < 10; i++ {
		tree1.insert(50*i, toggle)
		toggle = 1 - toggle
	}

	displayinorder(tree1.root)
}

func displayinorder(t *node) {
	if t == nil {
		return
	}
	displayinorder(t.left)
	fmt.Print(t.data, "-")
	displayinorder(t.right)
}
*/

package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(value int) {
	n := &Node{data: value}
	if t.root == nil {
		t.root = n
		return
	}
	curr := t.root

	for curr != nil {
		if n.data < curr.data {
			if curr.left == nil {
				curr.left = n
				return
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = n
				return
			}
			curr = curr.right
		}

	}

}

func displaypreorder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.data, "-")
	displaypreorder(n.left)
	displaypreorder(n.right)
}

func main() {
	test := &Tree{}
	test.insert(20)
	test.insert(30)
	test.insert(10)
	test.insert(25)
	test.insert(200)
	test.insert(17)

	displaypreorder(test.root)
}
