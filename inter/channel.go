package inter

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	Id        string    `json:"id,string"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"cretaed_at,string,omitempty"`
	Users     []string  `json:"users"`
}

func (channel *Channel) AddUser(id uuid.UUID) {
	channel.Users = append(channel.Users, id.String())
}

func (channel *Channel) NotifyAllUsers(notificationFunction func(string)) {
	for key := range channel.Users {

		notificationFunction(channel.Users[key])
	}
}
