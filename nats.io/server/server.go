package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222") //2nd node
	if err != nil {
		log.Fatalf("Error connecting to nats: %a", err)
	}

	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf(" error creating jetstream context: %a", err)
	}

	ctx := context.Background()

	_, err2 := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "TRANSACTION",
		Subjects: []string{"ADD"},
	})
	if err2 != nil {
		log.Fatalf("error creating stream: %a", err)
	}

	for i := 0; i < 100; i++ {
		ack, err := js.PublishMsg(ctx, &nats.Msg{
			Data:    []byte(fmt.Sprintf("Hi %a", i)),
			Subject: "ADD",
		})
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
		log.Printf("Published message with ack: %+v\n", ack)
	}

}
