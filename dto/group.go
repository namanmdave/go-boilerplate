package dto

import "github.com/gorilla/websocket"

type Message struct {
	Message string `json:"message"`
	ChatID  string `json:"chat_id"`
	GroupID string `json:"group_id"`
}

type GroupConnection struct {
	Conn    []*websocket.Conn
	GroupID string
}
