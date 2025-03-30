package directed_test

import (
	"testing"

	"github.com/sbiemont/grapo/directed"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCyclic(t *testing.T) {
	a := node{id: "a"}
	b := node{id: "b"}
	c := node{id: "c"}
	d := node{id: "d"}
	e := node{id: "e"}
	f := node{id: "f"}
	g := node{id: "g"}
	h := node{id: "h"}

	Convey("when ok #1", t, func() {
		Convey("when no edge (all independent nodes", func() {
			isCyclic := directed.IsCyclic[node](nil)
			So(isCyclic, ShouldBeFalse)
		})

		Convey("when mini graph", func() {
			dg := directed.Graph[node]{
				a: []node{c},
				b: []node{c},
			}
			isCyclic := directed.IsCyclic(dg)
			So(isCyclic, ShouldBeFalse)
		})

		Convey("when 2 graphs", func() {
			// a, b -> c
			// d -> e, f
			// f -> e, g, h
			dg := directed.Graph[node]{
				a: []node{c},
				b: []node{c},
				d: []node{e, f},
				f: []node{e, g, h},
			}
			isCyclic := directed.IsCyclic(dg)
			So(isCyclic, ShouldBeFalse)
		})

		// a -> c
		// c -> g
		// b -> c, e
		// d -> e, f
		// e -> g
		Convey("with order #1", func() {
			dg := directed.Graph[node]{
				a: []node{c},
				b: []node{c, e},
				c: []node{g},
				d: []node{e, f},
				e: []node{g},
			}
			isCyclic := directed.IsCyclic(dg)
			So(isCyclic, ShouldBeFalse)
		})

		Convey("with order #2", func() {
			dg := directed.Graph[node]{
				e: []node{g},
				d: []node{e, f},
				c: []node{g},
				b: []node{c, e},
				a: []node{c},
			}
			isCyclic := directed.IsCyclic(dg)
			So(isCyclic, ShouldBeFalse)
		})
	})

	Convey("when error", t, func() {
		Convey("custom #1", func() {
			// a -> b -> c -> d
			// a -> c
			// c -> a
			// d -> d
			dg := directed.Graph[node]{
				a: []node{b, c},
				b: []node{c},
				c: []node{d, a},
				d: []node{d},
			}
			isCyclic := directed.IsCyclic(dg)
			So(isCyclic, ShouldBeTrue)
		})

		Convey("custom #2", func() {
			// a -> b -> c -> d -> a
			dg := directed.Graph[node]{
				a: []node{b},
				b: []node{c},
				c: []node{d},
				d: []node{a},
			}
			isCyclic := directed.IsCyclic(dg)
			So(isCyclic, ShouldBeTrue)
		})
	})
}
