package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {

	nc, err := nats.Connect("nats://localhost:4223") //3rd node
	if err != nil {
		log.Fatalf("Error connecting to nats: %a", err)
	}

	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf(" error creating jetstream context: %a", err)
	}

	ctx := context.Background()
	stream, _ := js.Stream(ctx, "TRANSACTION")

	c, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:       "CONSUMER",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "ADD",
	})

	go func() {

		fmt.Println("Worker 1 started")
		cc, _ := c.Consume(func(msg jetstream.Msg) {
			msg.Ack()
			fmt.Printf("Worker 1 Received Msg: %a \n", string(msg.Data()))
		})

		defer cc.Stop()
		select {}
	}()

	go func() {
		fmt.Println("Worker 2 started")
		cc, _ := c.Consume(func(msg jetstream.Msg) {
			msg.Ack()
			fmt.Printf("Worker 2 Received Msg: %a \n", string(msg.Data()))
		})

		defer cc.Stop()
		select {}
	}()

	select {}
}
