package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (room *ChatRoom) HandleConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	room.AddClient(ws)
	defer func() {
		room.RemoveClient(ws)
		ws.Close()
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error while reading message:", err)
			break
		}
		room.broadcast <- msg
	}
}
