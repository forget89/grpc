package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "/phase/phase_proto"
)

type PhaseEqualibriumServer struct {
	pb.UnimplementedPhaseEqualibriumServer
}

func (s *PhaseEqualibriumServer) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.Response, error) {
	result := req.Num1 / req.Num2
	return &pb.Response{Response: result}, nil
}

func (s *PhaseEqualibriumServer) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.Response, error) {
	result := req.Num1 * req.Num2
	return &pb.Response{Response: result}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPhaseEqualibriumServer(s, &PhaseEqualibriumServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
