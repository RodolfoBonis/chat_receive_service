package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type MessageModel struct {
	Id       uuid.UUID `json:"id"`
	Sender   uuid.UUID `json:"sender"`
	Receiver uuid.UUID `json:"receiver"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
}

func (m *MessageModel) ToJSON() []byte {
	message, err := json.Marshal(m)

	if err != nil {
		log.Fatalf("Error marshalling message: %v", err)
	}

	return message
}
