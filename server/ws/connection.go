package ws

import (
	"go-boilerplate/dto"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var GroupClients = make(map[string]dto.GroupConnection)
var Broadcast = make(chan dto.Message, 100)
