// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/killtheverse/go-chat-signal-server/app/injectors"
)

// Injectors from wire.go:

func NewApp(config_path string) (*restApiApplication, error) {
	config, err := injectors.ProvideConfig(config_path)
	if err != nil {
		return nil, err
	}
	server := injectors.ProvideHttpServer(config)
	appRestApiApplication := &restApiApplication{
		config:     config,
		httpServer: server,
	}
	return appRestApiApplication, nil
}