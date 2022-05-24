package main

import (
	"github.com/google/wire"
	"github.com/killtheverse/go-chat-signal-server/internal/core/ports"
	"github.com/killtheverse/go-chat-signal-server/internal/core/services"
	"github.com/killtheverse/go-chat-signal-server/internal/handlers/httphandlers"
	"github.com/killtheverse/go-chat-signal-server/internal/repositories"
)

func InitializeUserHandlers(mongoURI repositories.ConnectionURI, mongoDBName repositories.DatabaseName) (*httphandlers.UserHandler, error) {
    wire.Build(
        wire.Bind(new(ports.UserRepository), new(*repositories.UserRepository)),
        wire.Bind(new(ports.UserService), new(*services.UserService)),
        repositories.NewUserRepository,
        services.NewUserService,
        httphandlers.NewUserHandler,
    )
    return &httphandlers.UserHandler{}, nil
}
