package main

import(
	stream "github.com/metinagaoglu/go-network-packet-listener/stream"
	"context"
);

func main() {
	ctx := context.Background()

	stream.Consume(ctx)
}
