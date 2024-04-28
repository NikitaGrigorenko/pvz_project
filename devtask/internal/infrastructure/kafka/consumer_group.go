package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
)

type MessageHandler interface {
	Handle(ctx context.Context, session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) error
}

type ConsumerGroupHandler struct {
	ready      chan bool
	msgHandler MessageHandler
}

func NewConsumerGroup(msgHandler MessageHandler) ConsumerGroupHandler {
	return ConsumerGroupHandler{
		ready:      make(chan bool),
		msgHandler: msgHandler,
	}
}

func (consumer *ConsumerGroupHandler) Ready() <-chan bool {
	return consumer.ready
}

func (consumer *ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			err := consumer.msgHandler.Handle(context.Background(), session, message)
			if err != nil {
				fmt.Println("Consumer group error", err)
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
