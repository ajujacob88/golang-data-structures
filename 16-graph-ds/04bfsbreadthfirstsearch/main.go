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

	test.BFS(1)
}

// BFS implements the breadth first search algorithm on the graph
func (g *Graph) BFS(start int) {
	// mark all vertices as unvisited
	visited := make(map[int]bool)
	for _, v := range g.vertices {
		visited[v.key] = false
	}

	// create a queue for BFS
	queue := []*Vertex{}

	// mark the start vertex as visited and enqueue it
	visited[start] = true
	queue = append(queue, g.getVertex(start))

	for len(queue) != 0 {
		// dequeue a vertex from queue and print it
		s := queue[0]
		queue = queue[1:]
		fmt.Printf("\n %v ", s.key)

		// get all adjacent vertices of the dequeued vertex s
		// if an adjacent vertex has not been visited, then mark it visited and enqueue it
		for _, v := range s.adjacent {
			if !visited[v.key] {
				visited[v.key] = true
				queue = append(queue, v)
			}
		}
	}
	fmt.Println("\n")
	fmt.Println(visited)

}
