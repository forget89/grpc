package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zibloidix/calc-api-grpc-golang/calc_api_pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client start")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := calc_api_pb.NewCalcServiceClient(cc)

	reqAdd := &calc_api_pb.AddRequest{
		A: 5,
		B: 5,
	}

	reqSubtract := &calc_api_pb.SubtractRequest{
		A: 500,
		B: 100,
	}

	reqMultiply := &calc_api_pb.MultiplyRequest{
		A: 5,
		B: 5,
	}

	reqDivide := &calc_api_pb.DivideRequest{
		A: 25,
		B: 5,
	}

	resAdd, err := c.Add(context.Background(), reqAdd)
	resSubtract, err := c.Subtract(context.Background(), reqSubtract)
	resMultiply, err := c.Multiply(context.Background(), reqMultiply)
	resDivide, err := c.Divide(context.Background(), reqDivide)

	log.Printf("Add: %v", resAdd)
	log.Printf("Subtract: %v", resSubtract)
	log.Printf("Multiply: %v", resMultiply)
	log.Printf("Divide: %v", resDivide)
}
