package main

import (
	"net"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Main", t, func() {

		Convey("main ", func() {
			myPatchs := NewPatches()
			defer myPatchs.Reset()

			go main()

			conn, err := net.Dial("tcp", ":3001")
			defer func() {
				if conn != nil {
					conn.Close()
				}
			}()

			So(err, ShouldBeNil)
		})

	})
}
