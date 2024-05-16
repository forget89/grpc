package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/forget89/grpc/phase"
)

func main() {
	serverAddress := "localhost:50051"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorServer(conn)

	divideResult, err := client.Divide(context.Background(), &pb.DivideRequest{Num1: 20.4, Num2: 2.4})
	if err != nil {
		log.Fatalf("Divide RPC failed: %v", err)
	}
	fmt.Printf("Divide Result: %.2f\n", divideResult.Response)

	multiplyResult, err := client.Divide(context.Background(), &pb.MultiplyRequest{Num1: 5.4, Num2: 2.4})
	if err != nil {
		log.Fatalf("Multiply RPC failed: %v", err)
	}
	fmt.Printf("Multiply Result: %.2f\n", multiplyResult.Response)
}
