package main

import "github.com/gorilla/websocket"

type ClientList map[*Client]bool

type Client struct {
	Conn *websocket.Conn
	manager *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		Conn: conn,
		manager: manager,
	}
}
