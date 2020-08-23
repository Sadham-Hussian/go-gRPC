package main

import (
	"context"
	"fmt"
	"time"

	sumallpb "github.com/Sadham-Hussian/go-gRPC/stream/client-streaming/sumAll/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := sumallpb.NewSumAllServiceClient(conn)

	stream, err := client.SumAll(context.Background())
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("sending %v into the stream\n", i)
		stream.Send(&sumallpb.NumberAddRequest{Num: int64(i)})
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	fmt.Println("Sum of Numbers: ", res.Result)
}
