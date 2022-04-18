package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/killtheverse/go-chat-signal-server/logging"
)

func ServeWs(w http.ResponseWriter, r *http.Request) {
    upgrader := websocket.Upgrader{}
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        logging.Write("[ERROR] Can't upgrade to websocket connection: %v", err)
    }
    logging.Write("%v", conn)
}
