package directed_test

import (
	"testing"

	"github.com/sbiemont/grapo/directed"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBFS(t *testing.T) {
	Convey("bfs", t, func() {
		a := node{id: "a"}
		b := node{id: "b"}
		c := node{id: "c"}
		d := node{id: "d"}
		e := node{id: "e"}
		f := node{id: "f"}
		g := node{id: "g"}
		h := node{id: "h"}

		Convey("when ok", func() {
			edges := directed.Graph[node]{
				a: {b, c, d},
				b: {e, f},
				c: {g, h},
			}
			var res []node
			err := directed.BFS(edges, a, func(n node) error {
				res = append(res, n)
				return nil
			})
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []node{a, b, c, d, e, f, g, h})
		})

		Convey("when cycle", func() {
			// a -> b <-> c
			edges := directed.Graph[node]{
				a: {b},
				b: {c},
				c: {b},
			}
			var res []node
			err := directed.BFS(edges, a, func(n node) error {
				res = append(res, n)
				return nil
			})
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []node{a, b, c})
		})

		Convey("when error", func() {
		})
	})
}
