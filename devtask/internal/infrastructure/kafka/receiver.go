package kafka

import (
	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"log"
	"time"
)

type Client struct {
	brokers []string
	Client  sarama.ConsumerGroup
}

func newConsumerGroup(brokers []string, group string) (sarama.ConsumerGroup, error) {
	log.Println("Starting a new Sarama consumer")

	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion

	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Session.Timeout = 60 * time.Second
	config.Consumer.Group.Rebalance.Timeout = 60 * time.Second
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	client, err := sarama.NewConsumerGroup(brokers, group, config)

	if err != nil {
		return nil, errors.Wrap(err, "error creating consumer group client")
	}

	return client, nil
}

func NewConsumer(brokers []string, group string) (*Client, error) {
	client, err := newConsumerGroup(brokers, group)

	if err != nil {
		return nil, err
	}

	consumer := &Client{
		brokers: brokers,
		Client:  client,
	}

	return consumer, nil
}
