package astar

// priorityQueue implements a priority queue for A* algorithm
type priorityQueue []*node

// Len returns the length of the priority queue
func (q priorityQueue) Len() int { return len(q) }

// Less compares two items in the queue
func (q priorityQueue) Less(i, j int) bool {
	return q[i].f < q[j].f
}

// Swap swaps two items in the priority queue
func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

// Push adds an item to the priority queue
func (q *priorityQueue) Push(x any) {
	item := x.(*node)
	item.index = len(*q)
	*q = append(*q, item)
}

// Pop removes and returns the item with the highest priority
func (q *priorityQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}
