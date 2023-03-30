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
	}
	curr := t.root

	for curr != nil {
		if n.data < curr.data {
			curr.left = n

			curr = curr.left
		} else {
			curr.right = n
			curr = curr.right
		}

	}

}

func (t *Tree) checkbst() bool {
	curr := t.root
	if t.root == nil {
		fmt.Println("no elements are present")
	}
	for curr != nil {

	}
	if curr.left.data <= curr.data {
		return true
	} else {
		return false
	}
}

/*

func (t *Tree) checkbst() bool {
	return t.bstimpl(math.MaxInt32, math.MinInt32)

}

func (n *Node) bstimpl(val1 int, val2 int) bool {
	curr:= n.data
	if root
}
*/

func main() {
	test := &Node{}
}
