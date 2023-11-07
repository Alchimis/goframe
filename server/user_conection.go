package server

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// в теории эту залупу можно будет сериализовать и что то добавить
type UserConnection struct {
	conn        *websocket.Conn
	connectedAt time.Time
	id          uuid.UUID
	userId      uuid.UUID
}

func ConstructUserConnection(conn *websocket.Conn) (*UserConnection, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &UserConnection{
		conn:        conn,
		connectedAt: time.Now(),
		id:          id,
	}, nil
}

func HandleConstructError(err error) {
	log.Println("error with Construct UserConnection", err)
}
