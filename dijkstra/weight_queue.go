package dijkstra

type path[T any] struct {
	weight float64 // total weight of the path
	nodes  []T     // list of nodes in the path
}

// weightQueue is a list of path ordered by total weight (node weight + distance)
// Implement heap.Interface for weightQueue[T]
type weightQueue[T any] []path[T]

func (q weightQueue[T]) Len() int           { return len(q) }
func (q weightQueue[T]) Less(i, j int) bool { return q[i].weight < q[j].weight }
func (q weightQueue[T]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *weightQueue[T]) Push(x any)        { *q = append(*q, x.(path[T])) }

func (q *weightQueue[T]) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}
