package main

import (
	"fmt"
	"io"
	"net"

	feedpb "github.com/Sadham-Hussian/go-gRPC/stream/bi-directional-streaming/feed/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	feedpb.RegisterFeedServer(srv, &server{})

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Broadcast(stream feedpb.Feed_BroadcastServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			panic(err)
		}

		feed := "New feed received: " + msg.GetFeed()
		fmt.Println("sending new feed..")
		stream.Send(&feedpb.FeedResponse{Feed: feed})
	}
}
