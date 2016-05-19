package hub

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/utrack/go-simple-chat/message"
	"strings"
	"testing"
	"time"
)

func TestHub(t *testing.T) {
	Convey("With hub", t, func() {
		h := NewHub(nil, nil, nil).(*hub)

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

		Convey("Client nickname sanitizer", func() {
			h.nameChecker = func(v string) (string, error) {
				if v == "should_fail" {
					return ``, errors.New("failed")
				}
				return "bar", nil
			}

			Convey("RegisterClient should fail with bad nick", func() {
				c1 := newClientMock()
				So(h.RegisterClient(c1, "should_fail"), ShouldNotBeNil)
			})
			Convey("Should accept new name from sanitizer", func() {
				c1 := newClientMock()
				So(h.RegisterClient(c1, "qwe"), ShouldBeNil)

				So(h.clientExists(`bar`), ShouldBeTrue)

				Convey("Should process the nick before checking for collisions", func() {
					So(h.RegisterClient(c1, `qwe`), ShouldEqual, ErrNickCollision)
				})
			})
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

			Convey("User should retrieve its own message", func() {
				msg := message.One{Type: message.EventMessage, IsMuted: true, Text: "barbaz"}
				c1.msgChan <- msg
				msg.From = cName

				var got message.One
				So(got, ShouldNotResemble, msg)
				select {
				case got = <-c1.acceptedMsgs:
				case <-time.After(time.Second):
					So(0, ShouldEqual, "client did not receive the message!")
				}
				So(got, ShouldResemble, msg)

				Convey("With sanitizer", func() {
					h.msgSanitizer = func(s string) string {
						return strings.Replace(s, `bar`, `foo`, -1)
					}

					Convey("Should sanitize the message's text", func() {
						msg := message.One{Type: message.EventMessage, IsMuted: true, Text: "barbaz"}
						c1.msgChan <- msg
						msg.From = cName
						msg.Text = `foobaz`

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

			Convey("With second client", func() {
				c2Name := "client2"
				c2 := newClientMock()
				So(h.RegisterClient(c2, c2Name), ShouldBeNil)

				Convey("Should receive own EventJoin", func() {
					// Flush EventJoin
					var got message.One
				ForOuter:
					for {
						select {
						case got = <-c2.acceptedMsgs:
							if got.From == c2Name {
								break ForOuter
							}
						case <-time.After(time.Second):
							So(0, ShouldEqual, "client did not receive the EventJoin!")
							break ForOuter
						}
					}
					So(got.Type, ShouldEqual, message.EventJoin)
					So(got.IsMuted, ShouldBeFalse)
					So(got.From, ShouldEqual, c2Name)

				})

				Convey("First one should receive the EventJoin", func() {
					var got message.One
					select {
					case got = <-c1.acceptedMsgs:
					case <-time.After(time.Second):
						So(0, ShouldEqual, "client did not receive the message!")
					}
					So(got.Type, ShouldEqual, message.EventJoin)
					So(got.From, ShouldEqual, c2Name)

					Convey("And nothing more", func() {
						var got message.One
						select {
						case got = <-c1.acceptedMsgs:
							So(got, ShouldEqual, "client received some message!")
						case <-time.After(time.Second):
						}
						So(got, ShouldResemble, message.One{})
					})
				})

				Convey("Second client should rcv muted EventJoins about existing users", func() {

					var got message.One
					select {
					case got = <-c2.acceptedMsgs:
					case <-time.After(time.Second):
						So(0, ShouldEqual, "client did not receive the EventJoin!")
					}
					So(got.Type, ShouldEqual, message.EventJoin)
					So(got.IsMuted, ShouldBeTrue)
					So(got.From, ShouldEqual, cName)
				})

				Convey("With EventJoins dumped", func() {
					select {
					case <-c1.acceptedMsgs:
					case <-time.After(time.Second):
					}

				ForOuter:
					for {
						select {
						case got = <-c2.acceptedMsgs:
						case <-time.After(time.Second / 2):
							break ForOuter
						}
					}

					Convey("First should receive EventLeave on second discon", func() {
						c2.discChan <- errors.New("test")

						var got message.One
						select {
						case got = <-c1.acceptedMsgs:
						case <-time.After(time.Second):
							So(0, ShouldEqual, "client did not receive the message!")
						}
						So(got.Type, ShouldEqual, message.EventLeave)
						So(got.From, ShouldEqual, c2Name)
						So(got.IsMuted, ShouldBeFalse)
						So(got.Text, ShouldEqual, "test")
					})

					Convey("Second should receive EventLeave on first discon", func() {
						c1.discChan <- errors.New("test")

						var got message.One
						select {
						case got = <-c2.acceptedMsgs:
						case <-time.After(time.Second):
							So(0, ShouldEqual, "client did not receive the message!")
						}
						So(got.Type, ShouldEqual, message.EventLeave)
						So(got.From, ShouldEqual, cName)
						So(got.IsMuted, ShouldBeFalse)
						So(got.Text, ShouldEqual, "test")
					})

					Convey("Clients should be able to chat", func() {
						Convey("First to second", func() {
							msg := message.One{Type: message.EventMessage, IsMuted: true, Text: "barbaz"}
							c1.msgChan <- msg
							msg.From = cName

							var got message.One
							So(got, ShouldNotResemble, msg)
							select {
							case got = <-c2.acceptedMsgs:
							case <-time.After(time.Second):
								So(0, ShouldEqual, "client did not receive the message!")
							}
							So(got, ShouldResemble, msg)
						})

						Convey("Second to first", func() {
							msg := message.One{Type: message.EventMessage, IsMuted: true, Text: "barbaz"}
							c2.msgChan <- msg
							msg.From = c2Name

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
			})
		})
	})
}
