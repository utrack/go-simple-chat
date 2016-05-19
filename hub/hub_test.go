package hub

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/utrack/go-simple-chat/message"
	"testing"
	"time"
)

func TestHub(t *testing.T) {
	Convey("With hub", t, func() {
		h := NewHub().(*hub)

		// Run the pumps
		h.Run()

		Convey("Chans and structs should be initialized", func() {
			So(h.incomingDiscons, ShouldNotBeNil)
			So(h.incomingMsgs, ShouldNotBeNil)
			So(h.sessions, ShouldNotBeNil)
		})

		Convey("ClientExists should return false", func() {
			So(h.clientExists("foo bar"), ShouldBeFalse)
		})

		Convey("Client should register successfully", func() {
			cName := "client1"
			c1 := newClientMock()
			So(h.RegisterClient(c1, cName), ShouldBeNil)

			// Flush EventJoin
			var got message.One
			select {
			case got = <-c1.acceptedMsgs:
			case <-time.After(time.Second):
				So(0, ShouldEqual, "client did not receive the EventJoin!")
			}
			So(got.Type, ShouldEqual, message.EventJoin)
			So(got.IsMuted, ShouldBeFalse)
			So(got.From, ShouldEqual, cName)

			Convey("ClientExists should return true", func() {
				So(h.clientExists(cName), ShouldBeTrue)
			})

			Convey("Double registering same nickname should fail", func() {
				So(h.RegisterClient(c1, cName), ShouldEqual, ErrNickCollision)
			})

			Convey("Client should be in the dict", func() {
				got, ok := h.sessions[cName]
				So(ok, ShouldBeTrue)
				So(got.c, ShouldResemble, c1)
			})

			Convey("Should deregister client on discon", func() {
				// TODO check logging
				c1.discChan <- errors.New("test")
				<-time.After(time.Second / 2)
				So(h.clientExists(cName), ShouldBeFalse)
				_, ok := h.sessions[cName]
				So(ok, ShouldBeFalse)
			})

			Convey("Should route the message to the client", func() {
				msg := message.One{Type: message.EventMessage, IsMuted: true, From: "foo", Text: "barbaz"}
				h.incomingMsgs <- msg

				var got message.One
				So(got, ShouldNotResemble, msg)
				select {
				case got = <-c1.acceptedMsgs:
				case <-time.After(time.Second):
					So(0, ShouldEqual, "client did not receive the message!")
				}
				So(got, ShouldResemble, msg)
			})

		})
	})
}
