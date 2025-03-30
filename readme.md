# grapo

`grapo` offers algorithms for graphs

algo | description
---- | -----------
`A*`  | A star algorithm to find the shortest path
`BFS` | Breadth first search
`DFS` | Depth-first search
`IsCyclic` | Detects cycles in a graph
`TopologicalSort` | Flattens a graph using topological sort

## Nodes definition

All algorithm can be used with your own comparable node definition

```golang
// example of a definition of a comparable node
type node struct {
  id string
}
```

If you want, you can use the provided directed graph definition `directed.Graph` (`map[T][]T`: a comparable node linked to an unordered list of nodes)

## A*

Generic `A*` algorithm.
See tests for a matrix or a custom graph application.

```golang
// Get an order list of nodes that represents the shortest path
path := astar.Run[node](
  start,                                    // start node
  goal,                                     // goal node
  weight func(node) float64 { .. },         // the weight of the node in parameter
  distance func(node, node) float64 { .. }, // the heuristic distance between the 2 given nodes
  neighbors func(node) []node { .. },       // list of neighbors of the node in parameter
)
```

Helper functions for heuristic distance:

* `astar.ManhattanDistance`
* `astar.EuclideanDistance`

## BFS (Breadth-first search)

Explore all nodes level by level starting with a given node

* The starting node is required
* Returns the first error raised
* Apply the process function each time a node is reached

```golang
a := node{id: "a"}
b := node{id: "b"}
c := node{id: "c"}
d := node{id: "d"}
e := node{id: "e"}

edges := map[node][]node{
  a: {b, c},
  b: {d, e},
  c: {e},
}
err := directed.BFS(edges, a, func(n node) error {
  fmt.Println(n) // or send in a slice, or in a channel, ...
  return nil
})
```

## DFS (Depth-first search)

Explore all nodes of the graph from the deepest level (the leaves) to the root(s)

* No starting node required
* Returns an error if a cycle is found in the graph or the first error raised
* Apply the process function each time a node is reached

```golang
// define your nodes
a := node{id: "a"}
// ...

// define your graph
edges := map[node][]node{
  // ...
}
err := directed.DFS(edges, func(n node) error {
  fmt.Println(n) // or send in a slice, or in a channel, ...
  return nil
})
```

## IsCyclic

Check if the graph has a cycle (it uses the DFS algorithm)

```golang
isCyclic := directed.IsCyclic(edges)
```

## TopologicalSort

Flattens the graph using a topological sort algorithm

* Returns an error if a cycle is found in the graph

```golang
flat, err := directed.TopologicalSort(edges)
```
