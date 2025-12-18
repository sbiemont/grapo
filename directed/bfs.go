package directed

// BFS performs a breadth-first search on a directed graph represented as an adjacency list
func BFS[T comparable](edges map[T][]T, first T, process func(T) error) error {
	// Start with the first node
	queue := []T{first}
	visited := make(map[T]bool)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue // Skip if already visited
		}

		// Process the node
		visited[node] = true
		err := process(node)
		if err != nil {
			return err
		}

		// Process the neighbors
		for _, neighbor := range edges[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return nil
}
