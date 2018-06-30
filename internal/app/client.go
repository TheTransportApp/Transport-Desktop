package app

import "fmt"

// Client is used to start the desktop app.
type Client struct {
	ConfigFile string
}

// Start will start the client to open the desktop window.
func (client *Client) Start() {
	fmt.Println("| Transport v1 starting ... |")
}
