package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "phase"
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

func (s *PhaseEqualibriumServer) Array(ctx context.Context, req *pb.ArrayRequest) (*pb.ArrayResponse, error) {
	return &pb.ArrayResponse{Array: req.Nums}, nil
}

func (s *PhaseEqualibriumServer) ArrayDivide(ctx context.Context, req *pb.ArrayDivideRequest) (*pb.ArrayDivideResponse, error) {
	result := make([]float64, len(req.Nums))
	for i, num := range req.Nums {
		result[i] = num / 100
	}
	log.Printf("ArrayDivide: %.4f\n", result)
	return &pb.ArrayDivideResponse{Array: result}, nil
}

func (s *PhaseEqualibriumServer) Fluid(ctx context.Context, req *pb.InitMessageRequest) (*pb.InitMessageResponse, error) {
	return &pb.InitMessageResponse{Fluid: req.fluids}, nil
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
