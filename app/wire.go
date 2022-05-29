//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/killtheverse/go-chat-signal-server/app/injectors"
)

func NewApp(config_path string) (*restApiApplication, error) {
    panic(wire.Build(
            injectors.ProvideConfig,
            injectors.ProvideHttpServer,
            wire.Struct(new(restApiApplication), "*"),
        ))
}
