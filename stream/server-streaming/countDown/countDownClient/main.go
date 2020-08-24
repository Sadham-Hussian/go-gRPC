package main

import (
	"context"
	"fmt"
	"io"

	countdownpb "github.com/Sadham-Hussian/go-gRPC/stream/server-streaming/countDown/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := countdownpb.NewCountdownClient(conn)

	timer := int64(10)

	stream, err := client.Start(context.Background(), &countdownpb.CountdownRequest{Timer: timer})
	if err != nil {
		panic(err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(msg)
	}
}
