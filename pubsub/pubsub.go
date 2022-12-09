package pubsub

import "context"

type Topic string

type Pubsub interface {
	Publish(ctx context.Context, channel Topic, data *Message) error
	Subscribe(ctx context.Context, channel Topic) (ch <-chan *Message, close func())
	//Unsubscribe(ctx context.Context, channel Channel) error
}
