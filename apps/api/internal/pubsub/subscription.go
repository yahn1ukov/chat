package pubsub

import "github.com/nats-io/nats.go"

type Subscription struct {
	sub  *nats.Subscription
	Data chan []byte
}

func (s *Subscription) Unsubscribe() {
	s.sub.Unsubscribe()
}
