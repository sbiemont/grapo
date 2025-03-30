package astar_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/sbiemont/grapo/astar"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAStarOnMatrix(t *testing.T) {
	// matrix is a 2D array of weights
	type matrix [][]float64

	// node is a 2D point
	type node struct {
		i, j int
	}

	// helper: builds a new node
	newNode := func(i, j int) node {
		return node{i: i, j: j}
	}

	// helper: builds a 6x6 matrix with a w weight
	// path to be found: follow the dots
	makeMatrix := func(w float64) matrix {
		return matrix{
			{0, 0, 0, 0, 0, 0}, // | . . . . . . |
			{w, w, w, w, w, 0}, // | W W W W W . |
			{0, 0, 0, 0, 0, 0}, // | . . . . . . |
			{0, w, w, w, w, w}, // | . W W W W W |
			{0, 0, 0, 0, 0, 0}, // | . . . . . . |
			{w, w, w, w, w, 0}, // | W W W W W . |
		}
	}

	// start and end nodes
	start := node{i: 0, j: 0}
	end := node{i: 5, j: 5}

	// weight of the node (get the value in the matrix of weigths)
	weight := func(m matrix, n node) float64 {
		return m[n.i][n.j]
	}

	// distance between 2 nodes in the matrix
	distance := func(a, b node) float64 {
		return astar.EuclideanDistance(float64(a.i), float64(a.j), float64(b.j), float64(b.j))
	}

	// neighbors of node in the matrix
	neighbors := func(m matrix, n node) []node {
		var nodes []node
		add := func(i, j int) {
			if i >= 0 && i < len(m) && j >= 0 && j < len(m[0]) {
				nodes = append(nodes, newNode(i, j))
			}
		}
		// 8 directions
		// up, down, left, right
		add(n.i-1, n.j)
		add(n.i+1, n.j)
		add(n.i, n.j-1)
		add(n.i, n.j+1)
		// diagonal up-left, up-right, down-left, down-right
		add(n.i-1, n.j-1)
		add(n.i-1, n.j+1)
		add(n.i+1, n.j-1)
		add(n.i+1, n.j+1)
		return nodes
	}

	Convey("when low weigths with shortcut", t, func() {
		m := makeMatrix(1)
		path := astar.Run(
			start, // start node
			end,   // end node
			func(n node) float64 { return weight(m, n) }, // weights
			distance, // distance between 2 nodes
			func(n node) []node { return neighbors(m, n) }, // neighbors of the given node
		)

		// | . . . . . . |      | o           |
		// | 1 1 1 1 1 . |      |   o         |
		// | . . . . . . |  =>  |     o       |
		// | . 1 1 1 1 1 |      |       o     |
		// | . . . . . . |      |         o   |
		// | 1 1 1 1 1 . |      |           o |
		So(path, ShouldResemble, []node{
			{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
		})
	})

	Convey("when meidum weigths with shortcut", t, func() {
		m := makeMatrix(4)
		path := astar.Run(
			start, // start node
			end,   // end node
			func(n node) float64 { return weight(m, n) }, // weights
			distance, // distance between 2 nodes
			func(n node) []node { return neighbors(m, n) }, // neighbors of the given node
		)

		// | . . . . . . |      | o o o o o   |
		// | 3 3 3 3 3 . |      |           o |
		// | . . . . . . |  =>  |           o |
		// | . 3 3 3 3 3 |      |           o |
		// | . . . . . . |      |           o |
		// | 3 3 3 3 3 . |      |           o |
		So(path, ShouldResemble, []node{
			{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
			{1, 5},
			{2, 5},
			{3, 5},
			{4, 5},
			{5, 5},
		})
	})

	Convey("when hight weigths without shortcut", t, func() {
		m := makeMatrix(9)
		path := astar.Run(
			start, // start node
			end,   // end node
			func(n node) float64 { return weight(m, n) }, // weights
			distance, // distance between 2 nodes
			func(n node) []node { return neighbors(m, n) }, // neighbors of the given node
		)

		fmt.Println(path)
		// | . . . . . . |      | o o o o o   |
		// | 9 9 9 9 9 . |      |           o |
		// | . . . . . . |  =>  |   o o o o   |
		// | . 9 9 9 9 9 |      | o           |
		// | . . . . . . |      |   o o o o   |
		// | 9 9 9 9 9 . |      |           o |
		So(path, ShouldResemble, []node{
			{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
			{1, 5},
			{2, 4}, {2, 3}, {2, 2}, {2, 1},
			{3, 0},
			{4, 1}, {4, 2}, {4, 3}, {4, 4},
			{5, 5},
		})
	})
}

func TestAStarOnNodeWithCoordinates(t *testing.T) {
	// node is a 2D point with coordinates, weight and neighbors
	type node struct {
		id        string
		x, y      float64
		weight    float64
		neighbors []*node
	}

	// distance = √((x2-x1)²-(y2-y1)²)
	distance := func(a, b *node) float64 {
		x := b.x - a.x
		y := b.y - a.y
		return math.Sqrt(x*x + y*y)
	}

	weight := func(a *node) float64 {
		return a.weight
	}

	neighbors := func(a *node) []*node {
		return a.neighbors
	}

	Convey("when ok", t, func() {
		// Create nodes
		nodeA := &node{id: "a"}
		nodeB := &node{id: "b"}
		nodeC := &node{id: "c"}
		nodeD := &node{id: "d"}
		nodeE := &node{id: "e"}

		// Create edges
		nodeA.neighbors = []*node{nodeB, nodeC}
		nodeB.neighbors = []*node{nodeC, nodeD}
		nodeC.neighbors = []*node{nodeD, nodeE}

		// Test A* algorithm
		path := astar.Run(nodeA, nodeE, weight, distance, neighbors)

		// Check the path
		So(path, ShouldResemble, []*node{nodeA, nodeC, nodeE})
	})

	Convey("when no path", t, func() {
		// Create nodes
		nodeA := &node{id: "a"}
		nodeB := &node{id: "b"}

		// Test A* algorithm with no path
		path := astar.Run(nodeA, nodeB, weight, distance, neighbors)
		So(path, ShouldBeNil)
	})

	Convey("when with weights", t, func() {
		nodeA := &node{id: "a", x: 4, y: 0, weight: 1}
		nodeB := &node{id: "b", x: 2, y: 3, weight: 1}
		nodeC := &node{id: "c", x: 0, y: 1, weight: 1}
		nodeD := &node{id: "d", x: 3, y: 2, weight: 100} // shortest path with massive weight
		nodeE := &node{id: "e", x: 4, y: 4, weight: 1}
		nodeF := &node{id: "f", x: 5, y: 3, weight: 1}
		nodeG := &node{id: "g", x: 5, y: 2, weight: 1}
		nodeH := &node{id: "h", x: 5, y: 1, weight: 1}
		nodeI := &node{id: "i", x: 6, y: 4, weight: 1}

		nodeA.neighbors = []*node{nodeB, nodeC, nodeD}
		nodeB.neighbors = []*node{nodeC, nodeF, nodeE}
		nodeC.neighbors = []*node{nodeB}
		nodeD.neighbors = []*node{nodeC, nodeG}
		nodeE.neighbors = []*node{nodeF}
		nodeF.neighbors = []*node{nodeI, nodeD}
		nodeI.neighbors = []*node{nodeG}
		nodeG.neighbors = []*node{nodeH}

		path := astar.Run(nodeA, nodeG, weight, distance, neighbors)
		So(path, ShouldResemble, []*node{nodeA, nodeB, nodeF, nodeI, nodeG})
	})
}
