package main

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	opts := []kgo.Opt{
		kgo.SeedBrokers("localhost:9092"),
		kgo.DefaultProduceTopic("my-topic"),
		kgo.ClientID("producer-client-id"),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		// TODO: handle/log this error
		return
	}
	defer client.Close()

	record := &kgo.Record{
		Value: []byte("Hello World"),
		Topic: "my-topic",
	}

	if err := client.ProduceSync(context.Background(), record).FirstErr(); err != nil {
		// TODO: handle/log this error
		return
	}
}
