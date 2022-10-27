package srv

import (
	"github.com/nats-io/nats.go"
)

func DefaultNatsServer() (*nats.EncodedConn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return nats.NewEncodedConn(nc, nats.JSON_ENCODER)
}

func NewNatsServer(address string, options ...nats.Option) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(address, options...)
	if err != nil {
		return nil, err
	}

	return nats.NewEncodedConn(nc, nats.JSON_ENCODER)

}
