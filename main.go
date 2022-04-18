package main

import (
	"github.com/killtheverse/go-chat-signal-server/app"
	"github.com/killtheverse/go-chat-signal-server/config"
)

func main() {
    config := config.NewConfig()
    app.ConfigurAppandRun(config)
}
