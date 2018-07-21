package web

import "github.com/graarh/golang-socketio"

type MessageHandler interface {
	GetChannel() string
	HandleMessage(handler func(c *gosocketio.Channel, msg interface{}))
}
