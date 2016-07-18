package main

// Search for a path from vs to vt and return an array that containes the
// string node of each intersection.
func DepthFirstSearch(g Graph, vs, vt string) (bool, []string) {

	// Initiate pred and color array with empty strings and default color WHITE.
	color := make(map[string]string, len(g))
	pred := make(map[string]string, len(g))
	for v := range g {
		pred[v] = ""
		color[v] = WHITE
	}

	// Run the visit and store the path from vs to vt in pred.
	dfsVisit(g, vs, vt, color, pred)

	// Construct a reverse path array from each node's pred.
	path := []string{}
	current_node := vt
	for current_node != "" {
		path = append([]string{current_node}, path...)
		current_node = pred[current_node]
	}
	// If the first item in path is vs, a path is found, otherwise, there is no
	// link between vs and vt. Thus, no path.
	foundPath := true
	if path[0] != vs {
		foundPath = false
	}

	return foundPath, path
}

// dfsVisit recursively visits all possible paths from branch u, until the
// destination t is reached. Branch is marked with color WHITE = unvisited,
// GRAY = processing and BLACK = visited.
func dfsVisit(g Graph, u, t string, color, pred map[string]string) {
	// Processing.
	color[u] = GRAY

	for _, v := range g[u] {
		// Recursive visit if branch has not been visited.
		if color[v] == WHITE {
			pred[v] = u
			// Only continue if target has not been reached.
			if v != t {
				dfsVisit(g, v, t, color, pred)
			}
		}
	}

	color[u] = BLACK
}
