package ui

import (
	"github.com/thetransportapp/transport-desktop/pkg/transport/ui/web"
)

type TransportUI struct {
	socketIo *web.SocketIOServer
}

func NewTransportUI() *TransportUI {
	c := &TransportUI{}
	return c
}

func (client *TransportUI) Run() {
	client.socketIo = web.NewSocketIOServer(1337)
	client.socketIo.Run()

	// Start UI application
}