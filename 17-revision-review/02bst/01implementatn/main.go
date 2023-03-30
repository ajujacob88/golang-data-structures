//for revision

package main

import (
	"fmt"
)

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

func (t *Tree) lookup(value int) {
	if t.root == nil {
		fmt.Println("not present")
	}
	curr := t.root
	for curr != nil {
		if value < curr.data {
			if value == curr.data {
				fmt.Println("present")
				return
			}
			curr = curr.left
		} else {
			if value == curr.data {
				fmt.Println("present")
				return
			}
			curr = curr.right
		}
	}
	fmt.Println("Not present")

}

func (t *Tree) delete(value int) {
	if t.root == nil {
		fmt.Println("Root is nil")
		return
	}
	curr := t.root
	if curr.right == nil && curr.left == nil {
		if curr.data == value {
			t.root = nil
		}
		return
	}

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
}

func (t *Tree) maxElement() {
	curr := t.root

	for curr.right != nil {
		if curr.right.right == nil {
			fmt.Println("max element is", curr.right.data)
			return
		}
		curr = curr.right

	}

}

func (t *Tree) minElement() {
	curr := t.root

	for curr.left != nil {
		if curr.left.left == nil {
			fmt.Println("max element is", curr.left.data)
			return
		}
		curr = curr.left

	}

}

func main() {
	test := &Tree{}
	test.insert(50)
	test.insert(55)
	test.insert(45)
	test.insert(20)
	test.insert(250)
	test.insert(10)
	test.insert(48)

	test.lookup(45)

	preorderprint(test.root)

	test.maxElement()
	test.minElement()

}

func preorderprint(n *Node) {
	if n == nil {
		return

	}

	fmt.Print(n.data, " - ")
	preorderprint(n.left)
	preorderprint(n.right)
}
