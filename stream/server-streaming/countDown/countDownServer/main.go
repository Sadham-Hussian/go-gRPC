package main

import (
	"net"
	"time"

	countdownpb "github.com/Sadham-Hussian/go-gRPC/stream/server-streaming/countDown/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	countdownpb.RegisterCountdownServer(srv, &server{})

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Start(req *countdownpb.CountdownRequest, stream countdownpb.Countdown_StartServer) error {
	timer := req.GetTimer()
	for timer > 0 {
		res := countdownpb.CountdownResponse{Count: timer}
		stream.Send(&res)
		timer = timer - 1
		time.Sleep(time.Second)
	}

	return nil
}
