package injectors

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/killtheverse/go-chat-signal-server/config"
)

func ProvideHttpServer(config *config.Config) (*http.Server) {
    httpServer := &http.Server{
        Addr : fmt.Sprintf(":%s", os.Getenv("PORT")),
        ReadTimeout: time.Second * time.Duration(config.Server.ReadTimeOut),
        WriteTimeout: time.Second * time.Duration(config.Server.WriteTimeOut),
        IdleTimeout: time.Second * time.Duration(config.Server.IdleTimeOut),
    }
    return httpServer
}
