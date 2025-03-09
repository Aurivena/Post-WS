package ws

import (
	"context"
	"log"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type PostMessage struct {
	Title       string
	Description string
	Author      string
	Date        time.Time
}

func write(ctx context.Context) {
	msg := PostMessage{
		Title:       "Тестовый заголовок",
		Description: "Тестовое описание",
		Author:      "Не nil",
		Date:        time.Now(),
	}

	for c := range connections {
		err := wsjson.Write(ctx, c, msg)
		if err != nil {
			log.Fatal("Error Write")
			delete(connections, c)
			return
		}
	}
}

func read(c *websocket.Conn, ctx context.Context) {
	for {
		var msg PostMessage
		err := wsjson.Read(ctx, c, &msg)
		if err != nil {
			log.Fatal("Error Read")
			break
		}
	}

}
