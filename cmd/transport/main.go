package main

import (
	"../../internal/app"
)

func main() {
	client := app.Client{ConfigFile: "Hi"}

	client.Start()
}
