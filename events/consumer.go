package events

type ConsumerFunc = func() (any, error)

type Consumer interface {
	Consume(string, any) (any, error)
}
