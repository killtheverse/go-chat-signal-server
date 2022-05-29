package main

import (
	"flag"

	"github.com/killtheverse/go-chat-signal-server/app"
)

func main() {
    var configFilePath string
    flag.StringVar(&configFilePath, "config", "config_files", "Absolute path to the configuration files directory")
    flag.Parse()

    application, err := app.NewApp(configFilePath)
    if err != nil {
        panic(err)
    }

    err = application.Init()
    if err != nil {
        panic(err)
    }

    application.StartHttpServer()
}
