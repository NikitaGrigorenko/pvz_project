//go:generate mockgen -source ./service.go -destination=./mocks/service.go -package=mock_service
package event

import (
	"devtask/internal/infrastructure/kafka"
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
)

type Producer interface {
	SendSyncMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error)
}

type KafkaSender struct {
	producer Producer
	topic    string
}

func NewProducer(brokers []string) (*kafka.Producer, error) {
	kafkaProducer, err := kafka.NewProducer(brokers)
	if err != nil {
		return nil, err
	}
	return kafkaProducer, nil
}

func NewKafkaSender(producer Producer, topic string) (*KafkaSender, error) {
	return &KafkaSender{
		producer: producer,
		topic:    topic,
	}, nil
}

func (s *KafkaSender) SendMessage(message model.EventMessage) error {
	kafkaMsg, err := s.buildMessage(message)
	if err != nil {
		fmt.Println("Send message marshal error", err)
		return err
	}

	_, _, err = s.producer.SendSyncMessage(kafkaMsg)

	if err != nil {
		return err
	}

	return nil
}

func (s *KafkaSender) buildMessage(message model.EventMessage) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Send message marshal error", err)
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
		Key:       sarama.StringEncoder(fmt.Sprint(message.Method)),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("test-header"),
				Value: []byte("test-value"),
			},
		},
	}, nil
}

func (s *KafkaSender) SendEvent(event model.EventMessage) error {
	err := s.SendMessage(model.EventMessage{
		Method:    event.Method,
		Body:      event.Body,
		TimeStamp: event.TimeStamp,
	})

	if err != nil {
		return errors.New("send sync message error")
	}
	return nil
}
