package events

import (
	"time"

	"github.com/nats-io/nats.go"
)

type NatsHandler = func(m *nats.Msg)

type Nats struct {
	nc *nats.EncodedConn
}

func NatsEvents(nc *nats.EncodedConn) *Nats {
	return &Nats{nc}
}

func (n *Nats) Consume(topic string, handler NatsHandler) (*nats.Subscription, error) {
	return n.nc.Subscribe(topic, handler)
}

func (n *Nats) Produce(topic string, msg any) error {
	return n.nc.Publish(topic, msg)
}

func (n *Nats) Request(topic string, payload any, response *any, timeout time.Duration) error {
	return n.nc.Request(topic, payload, &response, timeout)
}
