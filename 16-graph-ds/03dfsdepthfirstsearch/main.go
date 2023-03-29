//https://www.youtube.com/watch?v=vf-cxgUXcMk
//https://levelup.gitconnected.com/depth-breadth-first-search-in-go-8a6ddcdc73d9
//using chatgpt

package main

import (
	"fmt"
)

// graph strucutre - graph represents an adjacency list graph.
type Graph struct {
	vertices []*Vertex
}

// vertex structure - vertex represents a graph vertex
type Vertex struct {
	key int
	//because we are implementing adjacency list, each vertex in the graph will have a list of vertices that are holding the vertices of the neighbours.
	adjacent []*Vertex
}

// Add vertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
		//fmt.Println(g.vertices)
	}
}

// Add edge adds an edge to the graph - here we are doing the directed graph..
func (g *Graph) AddEdge(from, to int) {
	//get the address of vertices based on the keys
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid Edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing Edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// get vertex returns a pointer to the vertex with a keyinteger
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// print will print the adjacent list for each vertex of the graph. to print the actual keys or values to the keys..
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	fmt.Println(test)
	test.AddVertex(5)
	test.AddVertex(6)
	test.AddVertex(7)

	test.AddVertex(7)

	//test.Print()
	test.AddEdge(1, 2)
	//test.Print()
	test.AddEdge(9, 2)
	test.AddEdge(3, 7)
	test.AddEdge(3, 2)
	test.Print()

	//test.DFS2()
	test.DepthFirstSearch(1)

	g := &Graph{}
	for i := 0; i <= 5; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	fmt.Println("Depth-First Search:")
	g.DepthFirstSearch(1)
	//g.DFS2()
}

// here's an implementation of depth-first search (DFS) using a recursive approach. We start from a source vertex and traverse as deep as possible before backtracking.
// dfs implements the depth-first search algorithm
func (g *Graph) DFS(v *Vertex, visited map[int]bool) {
	// mark the vertex as visited
	visited[v.key] = true

	fmt.Printf("\n%v ", v.key)

	// recursively visit the adjacent vertices
	for _, n := range v.adjacent {
		if !visited[n.key] {
			g.DFS(n, visited)
		}
	}
}

// DepthFirstSearch implements the public API for depth-first search
func (g *Graph) DepthFirstSearch(start int) {
	visited := make(map[int]bool)
	v := g.getVertex(start)
	g.DFS(v, visited)
}

/*
// other program - not correct, wrong i think, so no need
func (g *Graph) DFS2() {
	var visited []int
	for _, vertex := range g.vertices {
		visited = vertex.DepthFirstSearch(visited)
	}
	fmt.Println(visited)
}

func (n *Vertex) DepthFirstSearch(array []int) []int {
	array = append(array, n.key)
	for _, child := range n.adjacent {
		array = child.DepthFirstSearch(array)
	}
	return array
}
*/
