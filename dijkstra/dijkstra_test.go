package dijkstra_test

import (
	"testing"

	"github.com/sbiemont/grapo/dijkstra"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDijkstra(t *testing.T) {
	type node struct {
		id        string
		weight    float64
		neighbors map[*node]float64
	}

	weight := func(a *node) float64 {
		return a.weight
	}

	neighbors := func(a *node) map[*node]float64 {
		return a.neighbors
	}

	Convey("when no weight", t, func() {
		// Create nodes
		nodeA := &node{id: "a"}
		nodeB := &node{id: "b"}
		nodeC := &node{id: "c"}
		nodeD := &node{id: "d"}
		nodeE := &node{id: "e"}

		// Create edges
		nodeA.neighbors = map[*node]float64{nodeB: 1, nodeC: 1}
		nodeB.neighbors = map[*node]float64{nodeC: 1, nodeD: 1}
		nodeC.neighbors = map[*node]float64{nodeD: 1, nodeE: 1}

		// Test algorithm
		path := dijkstra.Run(nodeA, nodeE, nil, neighbors)

		// Check the path
		So(path, ShouldResemble, []*node{nodeA, nodeC, nodeE})
	})

	Convey("when ok", t, func() {
		// Create nodes
		nodeA := &node{id: "a", weight: 6}
		nodeB := &node{id: "b", weight: 5}
		nodeC := &node{id: "c", weight: 4}
		nodeD := &node{id: "d", weight: 3}
		nodeE := &node{id: "e", weight: 2}
		nodeF := &node{id: "f", weight: 1}

		// Create edges
		nodeA.neighbors = map[*node]float64{nodeB: 5, nodeC: 2}
		nodeB.neighbors = map[*node]float64{nodeD: 8}
		nodeC.neighbors = map[*node]float64{nodeB: 7, nodeD: 4, nodeE: 8}
		nodeD.neighbors = map[*node]float64{nodeE: 6, nodeF: 4}
		nodeE.neighbors = map[*node]float64{nodeF: 3}

		// Test algorithm
		path := dijkstra.Run(nodeA, nodeF, weight, neighbors)

		// Check the path
		So(path, ShouldResemble, []*node{nodeA, nodeC, nodeD, nodeF})
	})

	Convey("when no path", t, func() {
		nodeA := &node{id: "a"}
		nodeB := &node{id: "b"}
		nodeC := &node{id: "c"}

		nodeA.neighbors = map[*node]float64{nodeB: 1}
		nodeB.neighbors = map[*node]float64{nodeA: 1}
		nodeC.neighbors = nil

		// Test algorithm
		path := dijkstra.Run(nodeA, nodeC, nil, neighbors)

		// Check the path
		So(path, ShouldBeNil)
	})
}
