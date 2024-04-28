package events

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type eventMessage struct {
	Method    string
	Body      string
	TimeStamp time.Time
}

type CustomMessageHandler struct{}

func (CustomMessageHandler) Handle(_ context.Context, _ sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) error {
	pm := eventMessage{}
	err := json.Unmarshal(message.Value, &pm)
	if err != nil {
		return err
	}

	log.Printf("Message claimed: value = %v, timestamp = %v, topic = %s",
		pm,
		message.Timestamp,
		message.Topic,
	)
	return nil
}
