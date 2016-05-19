/*Package clientBot provides simple pluggable bot example.
 */
package clientBot

import (
	"errors"
	"fmt"
	"github.com/utrack/go-simple-chat/client"
	"github.com/utrack/go-simple-chat/message"
	"time"
)

type bot struct {
	sendQueue  chan message.One
	disconChan chan error
}

// NewBot returns the new bot.
func NewBot() client.Client {
	ret := &bot{
		sendQueue:  make(chan message.One, 15),
		disconChan: make(chan error, 5),
	}
	go ret.pump()
	return ret
}

func (b *bot) Send(m message.One) error {
	if m.Type == message.EventJoin && !m.IsMuted {
		b.sendQueue <- message.One{
			Type: message.EventMessage,
			Text: fmt.Sprintf("Hello, %v!", m.From),
		}
	}
	return nil
}

func (b *bot) pump() {
	var ch <-chan time.Time
	ch = time.After(time.Second * time.Duration(random(5, 60)))

	for {
		<-ch
		text, err := getQuote()
		if err != nil {
			text = fmt.Sprintf("Joke failed :( Reason: %v", err.Error())
		}
		b.sendQueue <- message.One{Type: message.EventMessage, Text: text}
		ch = time.After(time.Second * time.Duration(random(5, 60)))
	}

}

func (b *bot) Disconnect() {
	b.disconChan <- errors.New("Gone Fishin'")
}

func (b *bot) DisconChan() <-chan error {
	return b.disconChan
}

func (b *bot) MsgChan() <-chan message.One {
	return b.sendQueue
}
