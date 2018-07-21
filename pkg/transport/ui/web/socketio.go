package web


import (
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"net/http"
	"log"
)

type SocketIOServer struct {
	server *gosocketio.Server
	port int
	handlers []MessageHandler
}

func NewSocketIOServer(port int) *SocketIOServer {
	socketIOServer := &SocketIOServer{
		server: gosocketio.NewServer(transport.GetDefaultWebsocketTransport()),
	}
	socketIOServer.port = port

	return socketIOServer
}

func (server *SocketIOServer) Run() {
	server.listenToMessages()
	server.createHttpServer()
}

func (server *SocketIOServer) AddMessageHandler(handler MessageHandler) {
	server.handlers = append(server.handlers, handler)
}

func (server *SocketIOServer) listenToMessages() {
	server.server.On(gosocketio.OnConnection, func(c *gosocketio.Channel, args interface{}) {
		log.Println("New client connected, client id is ", c.Id())
	})
	for _, handler := range server.handlers {
		server.server.On(handler.GetChannel(), handler.HandleMessage)
	}
}

func (server *SocketIOServer) createHttpServer() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server.server)
	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(":1337", serveMux))
}