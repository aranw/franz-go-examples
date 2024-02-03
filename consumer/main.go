package main

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	opts := []kgo.Opt{
		kgo.SeedBrokers("localhost:9092"),
		kgo.ClientID("consumer-client-id"),
		kgo.ConsumerGroup("my-group-identifier"),
		kgo.ConsumeTopics("my-topic"),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		// TODO: handle/log this error
		return
	}
	defer client.Close()

	for {
		fetches := client.PollFetches(context.Background())
		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			// returned from polls so that users can notice and take action.
			panic(fmt.Sprint(errs))
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), "from an iterator!")
		}
	}
}
