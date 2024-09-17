package controllers

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	mutex     sync.Mutex
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (r *ChatRoom) Run() {
	for msg := range r.broadcast {
		r.mutex.Lock()
		for client := range r.clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Error while sending message:", err)
				client.Close()
				delete(r.clients, client)
			}
		}
		r.mutex.Unlock()
	}
}

func (r *ChatRoom) AddClient(client *websocket.Conn) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.clients[client] = true
}

func (r *ChatRoom) RemoveClient(client *websocket.Conn) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.clients, client)
}
