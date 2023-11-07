package inter

import "time"

type Message struct {
	SenderId  string    `json:"sender_id"`
	ChannelId string    `json:"channel_id"`
	Message   string    `json:"message"`
	SendedAt  time.Time `json:"sended_at"`
}

func CreateMessage(chanelId, senderId, message string) Message {
	return Message{
		SenderId:  senderId,
		ChannelId: chanelId,
		Message:   message,
		SendedAt:  time.Now(),
	}
}
