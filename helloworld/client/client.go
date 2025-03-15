package main

import (
	"context"

	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"

	"helloworld/pb"
)

func main() {
	tRPCClient := pb.NewGreeterClientProxy(client.WithTarget("ip://127.0.0.1:8000"))
	resp, err := tRPCClient.Hello(context.Background(), &pb.HelloRequest{
		Msg: "world",
	})
	if err != nil {
		log.Error(err)
	}

	log.Info(resp.Msg)
}
