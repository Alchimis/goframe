package server

import (
	"github.com/gorilla/websocket"
)

type Server struct {
	Db                interface{}
	ActiveConnections map[string]*UserConnection
	Upgrader          websocket.Upgrader
}
