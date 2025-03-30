package astar

// priorityQueue implements a priority queue for A* algorithm
type priorityQueue []*node

// Len returns the length of the priority queue
func (pq priorityQueue) Len() int { return len(pq) }

// Less compares two items in the queue
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

// Swap swaps two items in the priority queue
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an item to the priority queue
func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*node)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
