package inter

import (
	"time"

	"github.com/google/uuid"
)

type UuidJson struct {
	Id uuid.UUID `json:"id,string"`
}

type User struct {
	Nickname        string    `json:"nickname"`
	Password        string    `json:"password"`
	Id              string    `json:"id"`
	JoinAt          time.Time `json:"join_at,string,omitempty"`
	LiveConnections []string  `json:"live_connections,omitempty"`
}
