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

func main() {
	graph := &Graph{map[int][]int{}}

	graph.insert(3, 5, true)
	graph.insert(3, 4, true)
	graph.insert(5, 6, false)
	// fmt.Println(graph.adjList)
	graph.display()

}
