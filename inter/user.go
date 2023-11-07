package inter

import (
	"time"

	"github.com/google/uuid"
)

type UuidJson struct {
	Id uuid.UUID `json:"id,string"`
}

type User struct {
	Nickname        string    `json:"nickname" clover:"nickname"`
	Password        string    `json:"password" clover:"password"`
	Id              string    `json:"id" clover:"id"`
	JoinAt          time.Time `json:"join_at,string,omitempty" clover:"join_at,string,omitempty"`
	LiveConnections []string  `json:"live_connections,omitempty" clover:"live_connections,omitempty"`
}
