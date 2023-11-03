package server

import (
	"time"

	"github.com/google/uuid"
)

type UUIDjson struct {
	Id uuid.UUID `json:"uuid"`
}

type Channel struct {
	UsersIds  []UUIDjson `json:"uuids"`
	ChannelId uuid.UUID  `json:"channel_id"`
	Name      string     `json:"channel_name"`
	CreatedAt time.Time  `json:"created_at"`
}

func (channel *Channel) AddUser(id uuid.UUID) {
	channel.UsersIds = append(channel.UsersIds, UUIDjson{id})
}
