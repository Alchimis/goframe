package inter

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	Id        string    `json:"id,string" clover:"id,string"`
	Name      string    `json:"name" clover:"name"`
	CreatedAt time.Time `json:"cretaed_at,string,omitempty" clover:"cretaed_at,string,omitempty"`
	Users     []string  `json:"users" clover:"users"`
}

func (channel *Channel) AddUser(id uuid.UUID) {
	channel.Users = append(channel.Users, id.String())
}

func (channel *Channel) NotifyAllUsers(notificationFunction func(string)) {
	for key := range channel.Users {

		notificationFunction(channel.Users[key])
	}
}

func NewChannel(name string) *Channel {
	return &Channel{
		Id:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
		Users:     []string{},
	}
}
