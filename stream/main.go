package stream

import (
	"fmt"
	"context"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"github.com/Shopify/sarama"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// the topic and broker address are initialized as constants
const (
	topic         = "network-log"
	brokerAddress = "kafka:9092"
)

func Produce(ctx context.Context,message []byte) {

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Balancer: &kafka.LeastBytes{},
		BatchSize: 1,
		BatchTimeout: 10 * time.Millisecond,
	})

	// each kafka message has a key and value. The key is used
	// to decide which partition (and consequently, which broker)
	// the message gets published on
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte("key"),
		// create an arbitrary message payload for the value
		Value: sarama.ByteEncoder(message),
	})

	if err != nil {
		fmt.Println(err)
		panic("could not write message " + err.Error())
	}

}

func Consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
	})

	fmt.Println("Start consume")
	for {

		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
	//	fmt.Println(msg)

		packet := gopacket.NewPacket(msg.Value,  layers.LayerTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
				fmt.Println(packet.String())
		//fmt.Println("received: ", string(msg.Value))
	}
}