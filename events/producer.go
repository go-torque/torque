package events

import "context"

type Producer interface {
	Produce(ctx context.Context, topic string, message any)
}
