package main

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	feedpb "github.com/Sadham-Hussian/go-gRPC/stream/bi-directional-streaming/feed/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := feedpb.NewFeedClient(conn)

	stream, err := client.Broadcast(context.Background())
	if err != nil {
		panic(err)
	}

	waitc := make(chan struct{})

	go func() {
		for i := 1; i <= 5; i++ {
			feed := "This is feed number " + strconv.Itoa(i)
			if err := stream.Send(&feedpb.FeedRequest{Feed: feed}); err != nil {
				panic(err)
			}

			time.Sleep(time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			panic(err)
		}
	}()

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				close(waitc)
				panic(err)
			}

			fmt.Println("New Feed Received : ", msg.GetFeed())

		}
	}()

	<-waitc
}
