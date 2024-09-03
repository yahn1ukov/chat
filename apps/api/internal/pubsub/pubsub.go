package pubsub

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/yahn1ukov/chat/apps/api/internal/config"
	"go.uber.org/fx"
)

type PubSub struct {
	conn *nats.Conn
}

type Params struct {
	fx.In

	Config *config.Config
}

func New(p Params) (*PubSub, error) {
	conn, err := nats.Connect(p.Config.NATS.Addr)
	if err != nil {
		return nil, err
	}

	return &PubSub{
		conn: conn,
	}, nil
}

func (p *PubSub) Publish(subject string, data any) {
	bytes, _ := json.Marshal(data)

	p.conn.Publish(subject, bytes)
}

func (p *PubSub) Subscribe(subject string) *Subscription {
	data := make(chan []byte, 1)

	sub, _ := p.conn.Subscribe(
		subject,
		func(msg *nats.Msg) {
			data <- msg.Data
		},
	)

	return &Subscription{
		sub:  sub,
		Data: data,
	}
}
