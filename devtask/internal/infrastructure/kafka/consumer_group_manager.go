package kafka

import (
	"context"
	"devtask/internal/service/events"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func ConsumerGroupManager(client Client, topic string) {
	keepRunning := true
	ctx, cancel := context.WithCancel(context.Background())
	consumer := NewConsumerGroup(events.CustomMessageHandler{})

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Client.Consume(ctx, []string{topic}, &consumer); err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-consumer.Ready()
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		}
	}

	cancel()
	wg.Wait()

	if err := client.Client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}
