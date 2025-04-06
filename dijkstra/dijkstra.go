package dijkstra

import (
	"container/heap"
)

// Inspired from https://dev.to/douglasmakey/implementation-of-dijkstra-using-heap-in-go-6e3
// Run finds the shortest path from start to goal using Dijkstra's algorithm
// * start:     first node of the path
// * goal:      last node of the path
// * weight:    give the node's weight (can be nil to give all nodes a 0 weight)
// * neighbors: list of unordered neighbors of the given node with the distance
// Returns the path or nil if nothing is found
func Run[T comparable](start, goal T, weight func(T) float64, neighbors func(T) map[T]float64) []T {
	// Init a new heap with a path containing the start node
	wqueue := &weightQueue[T]{}
	heap.Init(wqueue)
	heap.Push(wqueue, path[T]{
		nodes:  []T{start},
		weight: 0,
	})
	visited := make(map[T]bool)

	// While the queue is not empty, pop the path with the lowest weight
	for wqueue.Len() > 0 {
		p := heap.Pop(wqueue).(path[T]) // path with the lowest weight (and remove it from the heap)
		node := p.nodes[len(p.nodes)-1] // last node of the path
		if visited[node] {
			continue
		}
		if node == goal {
			return p.nodes
		}

		// For each neighbor of the current node, create a new path with its total weight
		for n, dist := range neighbors(node) {
			if !visited[n] {
				var w float64
				if weight != nil {
					w = weight(n)
				}
				heap.Push(wqueue, path[T]{ // new path with an increased weight and a new node
					nodes:  append(p.nodes, n),
					weight: p.weight + w + dist,
				})
			}
		}

		visited[node] = true
	}

	return nil
}
