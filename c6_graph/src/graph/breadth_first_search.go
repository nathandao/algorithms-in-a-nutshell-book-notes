package main

func BreadthFirstSearch(g Graph, vs, vt string) (int, []string) {
	// pred store each vertice's predicate vertice.
	pred := make(map[string]string, len(g))
	// dist is the distance for the vertice from vs.
	dist := make(map[string]int, len(g))
	// color marks the state of travel.
	color := make(map[string]string, len(g))

	// Set initial values.
	for v := range g {
		pred[v] = ""
		dist[v] = -1
		color[v] = WHITE
	}
	color[vs] = GRAY
	dist[vs] = 0
	q := []string{vs}

	// Start visiting layers of closest vertices from vs.
	for len(q) > 0 {
		u := q[0]
		lastVertice := u
		for _, v := range g[u] {
			if color[v] == WHITE && lastVertice != vt {
				dist[v] = dist[u] + 1
				pred[v] = u
				color[v] = GRAY
				q = append(q, v)
				lastVertice = v
			}
		}
		q = q[1:]
		color[u] = BLACK
	}

	// distance between vt and vs.
	distance := dist[vt]

	// get path from vt to vs if the 2 points are connected and have a distance.
	path := []string{}
	if distance != -1 {
		currentVertice := vt
		for currentVertice != "" {
			// Prepend vertice to the path array
			path = append([]string{currentVertice}, path...)
			currentVertice = pred[currentVertice]
		}
	}

	return distance, path
}
