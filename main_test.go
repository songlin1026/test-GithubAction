package main

import (
	"net"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Main", t, func() {

		Convey("main ", func() {

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
