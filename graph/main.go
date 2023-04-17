package main

import (
	"fmt"
)

type Graph struct {
	adjList map[int][]int
}

func (g *Graph) addVertex(data int) {
	g.adjList[data] = []int{}
}

func (g *Graph) insert(vertex int, edge int, isBidirectional bool) {

	if _, ok := g.adjList[vertex]; !ok {
		g.addVertex(vertex)
	}
	if _, ok := g.adjList[edge]; !ok {
		g.addVertex(edge)
	}
	g.adjList[vertex] = append(g.adjList[vertex], edge)
	if isBidirectional {
		g.adjList[edge] = append(g.adjList[edge], vertex)
	}

}

func (g *Graph) display() {
	for vertex, edges := range g.adjList {
		fmt.Printf("%d :", vertex)
		for i, edge := range edges {
			fmt.Printf("%d", edge)
			if i != len(edges)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}

// Graph node struct
type GraphNode struct {
	val       int
	neighbors []*GraphNode
}

// Breadth-first search traversal
func BFS(node *GraphNode) {
	queue := []*GraphNode{node}
	visited := make(map[*GraphNode]bool)
	visited[node] = true
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		fmt.Println(curr.val)
		for _, neighbor := range curr.neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

// Depth-first search traversal
func DFS(node *GraphNode, visited map[*GraphNode]bool) {
	visited[node] = true
	fmt.Println(node.val)
	for _, neighbor := range node.neighbors {
		if !visited[neighbor] {
			DFS(neighbor, visited)
		}
	}
}

func main() {
	graph := &Graph{map[int][]int{}}

	graph.insert(3, 5, true)
	graph.insert(3, 4, true)
	graph.insert(5, 6, false)
	// fmt.Println(graph.adjList)
	graph.display()

}
