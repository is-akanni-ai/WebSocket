package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// mark as used to avoid "unused" lint error when the variable is referenced only in closures
// var _ = allowedOrigins

var websocketUpgrader = websocket.Upgrader{
	CheckOrigin:     CheckOrigin,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	clients ClientList
	sync.RWMutex

	// Clients    map[*websocket.Conn]bool
	// Broadcast  chan []byte
	// Register   chan *websocket.Conn
	// Unregister chan *websocket.Conn
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
		// 	Clients:    make(map[*websocket.Conn]bool),
		// 	Broadcast:  make(chan []byte),
		// 	Register:   make(chan *websocket.Conn),
		// 	Unregister: make(chan *websocket.Conn),
	}
}

func (m *Manager) serverWS(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket Endpoint Hit")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("websocket upgrade failed: %v", err)
		return
	}

	client := NewClient(conn, m)
	m.addClient(client)
}

func (m *Manager) addClient(Clients *Client) {
	m.Lock();
	defer m.Unlock()

	m.clients[client] = true;
}
