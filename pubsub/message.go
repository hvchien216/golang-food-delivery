package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	channel   Topic // can be ignored
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()
	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (event *Message) String() string {
	return fmt.Sprintf("Message %s", event.channel)
}

func (event *Message) Channel() Topic {
	return event.channel
}

func (event *Message) SetChannel(channel Topic) {
	event.channel = channel
}

func (event *Message) Data() interface{} {
	return event.data
}
