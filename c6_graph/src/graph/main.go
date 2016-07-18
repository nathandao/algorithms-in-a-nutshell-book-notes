package main

import "fmt"

const WHITE = "WHITE"
const BLACK = "BLACK"
const GRAY = "GRAY"

type Graph map[string][]string

func main() {

	// Use a simple map[string][]string to represent a graph of vertices and a
	// slice of each of their corresponding target vertices (edges).
	// For a visual representation, check the image ./graph.png.
	G := Graph{
		"A": []string{"B", "C", "D", "E", "F"},
		"B": []string{"A", "D", "G"},
		"C": []string{"A", "D"},
		"D": []string{"A", "B", "E"},
		"E": []string{"A", "D", "H"},
		"F": []string{"A", "G"},
		"G": []string{"B", "F", "H"},
		"H": []string{"G", "E"},
		"I": []string{},
	}

	fmt.Println("Undirected unweighted graph:")
	for k, v := range G {
		fmt.Println(k, ":", v)
	}
	fmt.Println("\n**NOTE: SEE IMAGE './graph.png' FOR VISUAL PRESENTATION OF THE GRAPH\n")

	fmt.Println("=== DEPTH FIRST SEARCH ===")

	found, path := DepthFirstSearch(G, "A", "H")
	fmt.Println("Path exists from A to H:", found, ", path:", path)

	found, path = DepthFirstSearch(G, "G", "B")
	fmt.Println("Path exists from G to B:", found, ", path:", path)

	found, path = DepthFirstSearch(G, "A", "I")
	fmt.Println("Path exists from A to I:", found, ", path:", path)

	fmt.Println("=== BREADTH FIRST SEARCH ===")

	dist, path := BreadthFirstSearch(G, "A", "H")
	fmt.Println("Distance from A to H:", dist, ", path:", path)

	dist, path = BreadthFirstSearch(G, "A", "B")
	fmt.Println("Distance from A to B:", dist, ", path:", path)

	dist, path = BreadthFirstSearch(G, "A", "I")
	fmt.Println("Distance from A to I:", dist, ", path:", path)
}
