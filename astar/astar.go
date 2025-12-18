package astar

import (
	"container/heap"
	"math"
	"slices"
)

// ManhattanDistance is the sum of the absolute differences of their coordinates
// |x1 - x2| + |y1 - y2|
// Helper function
func ManhattanDistance(x1, y1, x2, y2 float64) float64 {
	return math.Abs(x1-x2) + math.Abs(y1-y2)
}

// EuclideanDistance calculates the straight-line distance between two points in a 2D space
// √((x2 - x1)² + (y2 - y1)²)
// Helper function
func EuclideanDistance(x1, y1, x2, y2 float64) float64 {
	x := float64(x2 - x1)
	y := float64(y2 - y1)
	return math.Sqrt(x*x + y*y)
}

// node is an internal struct to store data
type node[T comparable] struct {
	value  T        // actual value of the node
	f      float64  // total estimated cost (f=g+h)
	g      float64  // cost from current node
	h      float64  // heuristic (estimated) distance from node to goal
	index  int      // for priority queue
	parent *node[T] // parent computed during A* search
}

// converter is a map to store nodes and their corresponding values
// it uses two maps to allow bidirectional lookup
type converter[T comparable] struct {
	cache map[T]*node[T]
}

// newConverter creates a new empty cache for the given type
func newConverter[T comparable]() *converter[T] {
	return &converter[T]{
		cache: make(map[T]*node[T]),
	}
}

// fetch retrieves the node from the cache or creates a new one if it doesn't exist
func (c converter[T]) fetch(n T) *node[T] {
	m, ok := c.cache[n]
	if ok {
		return m
	}
	m = &node[T]{value: n}
	c.cache[n] = m
	return m
}

// Run performs the A* search algorithm
// * start:     first node of the path
// * goal:      last node of the path
// * weight:    give the node's weight (can be nil to give all nodes a 0 weight)
// * distance:  heuristic (estimated) distance between 2 nodes
// * neighbors: list of unordered neighbors of the given node
// Returns the found path or nil if nothing is found
func Run[T comparable](start, goal T, weight func(T) float64, distance func(T, T) float64, neighbors func(T) []T) []T {
	// Initialize opened and closed lists
	c := newConverter[T]()
	startNode := c.fetch(start)
	openedList := map[*node[T]]struct{}{startNode: {}}
	closedList := map[*node[T]]struct{}{}
	queue := &priorityQueue[T]{}
	heap.Init(queue)
	heap.Push(queue, startNode)

	// Initialize node properties
	startNode.g = 0                         // Cost from start to start is 0
	startNode.h = distance(start, goal)     // Estimate to goal
	startNode.f = startNode.g + startNode.h // Total estimated cost
	startNode.parent = nil                  // For path reconstruction

	goalNode := c.fetch(goal)

	for len(openedList) > 0 {
		// The queue is empty: path found
		if queue.Len() == 0 {
			return nil
		}

		// Get node with lowest f value (from prority queue)
		currentNode := heap.Pop(queue).(*node[T])
		current := currentNode.value

		// Check if we've reached the goal
		if currentNode == goalNode {
			return path(startNode, currentNode)
		}
		// Move current node from opened to closed list
		delete(openedList, currentNode)
		closedList[currentNode] = struct{}{}

		// Check all neighboring nodes
		for _, neighbor := range neighbors(current) {
			neighborNode := c.fetch(neighbor)
			if _, ok := closedList[neighborNode]; ok {
				continue // Skip already evaluated nodes
			}

			// Estimage g
			gEstimated := currentNode.g
			if weight != nil {
				gEstimated += weight(current)
			}
			_, opened := openedList[neighborNode]
			switch {
			case !opened:
				openedList[neighborNode] = struct{}{} // add neighbor to opened list
			case gEstimated >= neighborNode.g:
				continue // bad path, next node
			default:
				heap.Remove(queue, neighborNode.index) // remove neighbor from priority queue
			}

			// Best current path
			neighborNode.parent = currentNode
			neighborNode.g = gEstimated
			neighborNode.h = distance(neighbor, goal)
			neighborNode.f = neighborNode.g + neighborNode.h
			heap.Push(queue, neighborNode)
		}
	}
	return nil // no path found
}

// path from start to current node
func path[T comparable](start, current *node[T]) []T {
	// run from current to start node using parents and inverse the path
	var path []*node[T]
	for current != start {
		path = append(path, current)
		current = current.parent
	}
	path = append(path, start)
	slices.Reverse(path)

	// convert using cache
	nodes := make([]T, len(path))
	for i, p := range path {
		nodes[i] = p.value
	}
	return nodes
}
