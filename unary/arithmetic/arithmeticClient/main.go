package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Sadham-Hussian/go-gRPC/unary/arithmetic/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	var choice string
	var a, b int
	for {
		fmt.Println("Enter add to Add")
		fmt.Println("Enter sub to Subtract")
		fmt.Println("Enter mul to Multiply")
		fmt.Println("Enter div to Divide")
		fmt.Println("Enter exit to Exit")
		fmt.Println("Enter ur choice: ")
		fmt.Scanln(&choice)
		fmt.Println("Enter a: ")
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		req := &proto.Request{A: int64(a), B: int64(b)}
		switch choice {
		case "add":
			if res, err := client.Add(context.Background(), req); err != nil {
				panic(err)
			} else {
				fmt.Println(res.Result)
			}
		case "sub":
			if res, err := client.Subtract(context.Background(), req); err != nil {
				panic(err)
			} else {
				fmt.Println(res.Result)
			}
		case "mul":
			if res, err := client.Multiply(context.Background(), req); err != nil {
				panic(err)
			} else {
				fmt.Println(res.Result)
			}
		case "div":
			if res, err := client.Divide(context.Background(), req); err != nil {
				panic(err)
			} else {
				fmt.Println(res.Result)
			}
		case "exit":
			os.Exit(1)
		default:
			fmt.Println("Enter valid option")
		}
	}
}
