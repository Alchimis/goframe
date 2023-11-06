package internal

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	Id        uuid.UUID `json:"id,string"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"cretaed_at,string,omitempty"`
}
