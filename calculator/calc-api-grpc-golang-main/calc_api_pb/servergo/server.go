package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/zibloidix/calc-api-grpc-golang/calc_api_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(ctx context.Context, req *calc_api_pb.AddRequest) (*calc_api_pb.CalcResponse, error) {
	a := req.GetA()
	b := req.GetB()
	res := &calc_api_pb.CalcResponse{
		Result: a + b,
	}
	return res, nil
}

func (*server) Subtract(ctx context.Context, req *calc_api_pb.SubtractRequest) (*calc_api_pb.CalcResponse, error) {
	a := req.GetA()
	b := req.GetB()
	res := &calc_api_pb.CalcResponse{
		Result: a - b,
	}
	return res, nil
}

func (*server) Multiply(ctx context.Context, req *calc_api_pb.MultiplyRequest) (*calc_api_pb.CalcResponse, error) {
	a := req.GetA()
	b := req.GetB()
	res := &calc_api_pb.CalcResponse{
		Result: a * b,
	}
	return res, nil
}

func (*server) Divide(ctx context.Context, req *calc_api_pb.DivideRequest) (*calc_api_pb.CalcResponse, error) {
	a := req.GetA()
	b := req.GetB()
	res := &calc_api_pb.CalcResponse{
		Result: a / b,
	}
	return res, nil
}

func main() {
	fmt.Println("Calc gRPC service start")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	s := grpc.NewServer()
	calc_api_pb.RegisterCalcServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
