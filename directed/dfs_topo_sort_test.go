package directed_test

import (
	"testing"

	"github.com/sbiemont/grapo/directed"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTopologicalSort(t *testing.T) {
	a := node{id: "a"}
	b := node{id: "b"}
	c := node{id: "c"}
	d := node{id: "d"}
	e := node{id: "e"}
	f := node{id: "f"}
	g := node{id: "g"}
	h := node{id: "h"}

	Convey("topological sort", t, func() {
		Convey("when no edge (all independent nodes", func() {
			dg := directed.Graph[node](nil)

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeNil)
			So(topo, ShouldBeEmpty)
		})

		Convey("when one way", func() {
			// a -> b -> c
			dg := directed.Graph[node]{
				a: []node{b},
				b: []node{c},
			}

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeNil)
			So(ShouldBeOrdered(topo, a, b, c), ShouldBeTrue)
		})

		Convey("when mini graph #1", func() {
			// a -> c
			// b -> c
			// Valid topo are
			// * a, b, c
			// * b, a, c
			dg := directed.Graph[node]{
				a: []node{c},
				b: []node{c},
			}

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeNil)
			So(ShouldBeOrdered(topo, a, c), ShouldBeTrue)
			So(ShouldBeOrdered(topo, b, c), ShouldBeTrue)
		})

		Convey("when mini graph #2", func() {
			// a -> b, c
			// Valid topo are
			// * a, b, c
			// * a, c, b
			dg := directed.Graph[node]{
				a: []node{b, c},
			}

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeNil)
			So(ShouldBeOrdered(topo, a, b), ShouldBeTrue)
			So(ShouldBeOrdered(topo, a, c), ShouldBeTrue)
		})

		Convey("when 2 graphs", func() {
			// d -> e, f
			// f -> e, g, h
			// Valid topo are
			// * d, e, f, g, h
			// * d, e, f, h, g
			// * d, f, e, g, h
			// * d, f, e, h, g
			dg := directed.Graph[node]{
				d: []node{e, f},
				f: []node{g, h},
			}

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeNil)
			So(ShouldBeOrdered(topo, d, e), ShouldBeTrue)
			So(ShouldBeOrdered(topo, d, f), ShouldBeTrue)
			So(ShouldBeOrdered(topo, f, g), ShouldBeTrue)
			So(ShouldBeOrdered(topo, f, h), ShouldBeTrue)
		})

		Convey("when topo", func() {
			dg := directed.Graph[node]{
				a: []node{b, f},
				b: []node{c},
				c: []node{d},
				e: []node{f},
			}

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeNil)
			So(ShouldBeOrdered(topo, a, b, c, d), ShouldBeTrue)
			So(ShouldBeOrdered(topo, a, f), ShouldBeTrue)
			So(ShouldBeOrdered(topo, e, f), ShouldBeTrue)
		})

		Convey("with order", func() {
			dg := directed.Graph[node]{
				a: []node{c},
				b: []node{c, e},
				c: []node{g},
				d: []node{e, f},
				e: []node{g},
			}

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeNil)
			So(ShouldBeOrdered(topo, a, c, g), ShouldBeTrue)
			So(ShouldBeOrdered(topo, b, c, g), ShouldBeTrue)
			So(ShouldBeOrdered(topo, b, e, g), ShouldBeTrue)
			So(ShouldBeOrdered(topo, d, e), ShouldBeTrue)
			So(ShouldBeOrdered(topo, d, f), ShouldBeTrue)
		})
	})

	Convey("when error", t, func() {
		Convey("custom #1", func() {
			// a -> b -> c -> d
			// a -> c -> a (loop)
			// d -> d
			dg := directed.Graph[node]{
				a: []node{b, c},
				b: []node{c},
				c: []node{d, a},
				d: []node{d},
			}

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeError, directed.ErrCyclicGraph.Error())
			So(topo, ShouldBeNil)
		})

		Convey("custom #2", func() {
			// a -> b -> c -> d -> a
			dg := directed.Graph[node]{
				a: []node{b},
				b: []node{c},
				c: []node{d},
				d: []node{a},
			}

			topo, err := directed.TopologicalSort(dg)
			So(err, ShouldBeError, directed.ErrCyclicGraph.Error())
			So(topo, ShouldBeNil)
		})
	})
}
