package main

import (
	"log"
	"net"

	pb "grpc-istio/pb"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMultServiceServer(s, &server {})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *pb.MultRequest) (*pb.MultResponse, error) {
	a, b := r.A, r.B
	return &pb.MultResponse{ Result: a*b }, nil
}

func (s *server) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	a, b, c := r.A, r.B, r.C
	return &pb.ValidateResponse{ Valid: a*b == c }, nil
}
