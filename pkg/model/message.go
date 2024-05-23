package model

import (
	"encoding/json"
	"time"
)

// Message is used as a packet to be sent over EMQX cluster
type Message struct {
	Sent time.Time `json:"sent"`
}

func (m Message) ToBytes() []byte {
	bytes, _ := json.Marshal(m)

	return bytes
}

func CreateMessageFromBytes(bytes []byte) *Message {
	instance := &Message{}

	if err := json.Unmarshal(bytes, instance); err != nil {
		return nil
	}

	return instance
}
