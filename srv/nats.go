package srv

import (
	"github.com/nats-io/nats.go"
)

func NewNatsServer(address string, options ...nats.Option) (*nats.Conn, error) {
	return nats.Connect(address, options...)
}
