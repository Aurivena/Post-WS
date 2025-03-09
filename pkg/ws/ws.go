package ws

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
)

var (
	connections = make(map[*websocket.Conn]bool)
	mutex       sync.Mutex
)

func WsConnect(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println("Error accept")
		return
	}

	mutex.Lock()
	connections[c] = true
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(connections, c)
		mutex.Unlock()
		c.CloseNow()
	}()

	ctx := context.Background()
	write(ctx)
	read(c, ctx)
}
