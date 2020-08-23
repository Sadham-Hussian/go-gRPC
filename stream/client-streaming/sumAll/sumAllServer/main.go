package main

import (
	"io"
	"net"

	sumallpb "github.com/Sadham-Hussian/go-gRPC/stream/client-streaming/sumAll/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	sumallpb.RegisterSumAllServiceServer(srv, &server{})

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

// SumAll sums up all the numbers send by the client and returns the sum
func (s *server) SumAll(stream sumallpb.SumAllService_SumAllServer) error {
	var sum int64

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&sumallpb.SumResponse{Result: sum})
		}

		if err != nil {
			panic(err)
		}

		sum = sum + msg.GetNum()
	}
}
