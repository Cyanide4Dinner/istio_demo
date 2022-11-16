package main

import (
	"fmt"
	"context"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc_client/pb"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("192.168.49.2:31732", opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	md := metadata.New(map[string]string{ "actor_type": "retail" })

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// err = grpc.SendHeader(ctx, md)
	// if err != nil {
	// 	log.Fatalf("Error sending header: %v", err)
	// }

	client := pb.NewMultServiceClient(conn)
	res, err := client.Compute(ctx, &pb.MultRequest{ A: 7, B: 9 })
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	fmt.Println(res)
}
