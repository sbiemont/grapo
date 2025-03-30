package directed

// Graph represents a directed graph using an adjacency list
// This type is an exemple, algorithms can be used without it, just using a map[T][]T
type Graph[T comparable] map[T][]T
