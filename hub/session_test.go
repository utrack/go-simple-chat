package hub

import (
	"github.com/utrack/go-simple-chat/message"

	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

// mock of client.Client.
type clientMock struct {
	// messages from the hub
	acceptedMsgs chan message.One
	// discon chan, returned to the Session
	discChan chan error
	// msgChan, messages rcvd by the client
	msgChan chan message.One
}

func (c *clientMock) Send(m message.One) error {
	c.acceptedMsgs <- m
	return nil
}

func (c *clientMock) DisconChan() <-chan error {
	return c.discChan
}

func (c *clientMock) MsgChan() <-chan message.One {
	return c.msgChan
}

func newClientMock() *clientMock {
	return &clientMock{
		acceptedMsgs: make(chan message.One, 30),
		discChan:     make(chan error, 5),
		msgChan:      make(chan message.One, 30),
	}

}

func TestSession(t *testing.T) {
	Convey("With setup", t, func() {
		// mock hub's channels
		msgChan := make(chan message.One, 30)
		disconChan := make(chan disconMsg, 30)

		// Mock client
		c := newClientMock()

		name := `SomeName`
		// Session to test
		sess := newSession(c, name, msgChan, disconChan)
		go sess.runPump()

		Convey("Should relay messages to the hub", func() {
			msg := message.One{Type: message.EventMessage, IsMuted: true, From: "foobar", Text: "hello world"}
			c.msgChan <- msg

			var got message.One

			select {
			case got = <-msgChan:
			case <-time.After(time.Second):
				So(0, ShouldEqual, "Hub didn't receive the message!")
			}

			Convey("Should block bad/unknown EventTypes", func() {
				for i := 0; i < 10; i++ {
					if i == message.EventMessage || i == message.EventPresenceState {
						continue
					}

					msg := message.One{Type: message.EventType(i), IsMuted: true, From: "foobar", Text: "hello world"}
					c.msgChan <- msg

					var got message.One
					select {
					case got = <-msgChan:
						So(got, ShouldEqual, "Hub shouldn't receive the message!")
					case <-time.After(time.Millisecond * 200):
					}
				}
			})
			Convey("Message's data should be equal to the sent one", func() {
				So(got.Text, ShouldEqual, msg.Text)
				So(got.IsMuted, ShouldEqual, msg.IsMuted)
				So(got.Type, ShouldEqual, msg.Type)
			})
			Convey("Name should be changed to the session's name", func() {
				So(got.From, ShouldEqual, name)
				So(got.From, ShouldNotEqual, msg.From)
			})

		})

		Convey("Should relay messages from the hub", func() {
			msg := message.One{From: "fromfield", Type: message.EventPresenceState, IsMuted: true, Text: "qwww"}
			err := sess.send(msg)

			So(err, ShouldBeNil)
			var got message.One
			select {
			case got = <-c.acceptedMsgs:
			case <-time.After(time.Second):
				So(0, ShouldEqual, "Client didn't receive the message!")
			}
			So(got, ShouldResemble, msg)
		})

		Convey("On disconnection", func() {
			reason := errors.New("Spite")
			c.discChan <- reason

			Convey("Should create EventLeave", func() {
				var got message.One
				select {
				case got = <-msgChan:
				case <-time.After(time.Second):
					So(0, ShouldEqual, "Hub didn't received EventLeave!")
				}

				So(got.Type, ShouldEqual, message.EventLeave)
				So(got.From, ShouldEqual, name)
				So(got.Text, ShouldEqual, reason.Error())
			})

			Convey("Should notify the hub", func() {
				var got disconMsg
				select {
				case got = <-disconChan:
				case <-time.After(time.Second):
					So(0, ShouldEqual, "Hub didn't receive disconMsg!")
				}
				So(got.reason, ShouldEqual, reason.Error())
				So(got.name, ShouldEqual, name)
			})
		})
	})
}
