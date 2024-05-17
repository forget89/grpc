package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "phase"
)

func main() {
	serverAddress := "localhost:50051"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}

	defer conn.Close()

	client := pb.NewPhaseEqualibriumClient(conn)

	divideResult, err := client.Divide(context.Background(), &pb.DivideRequest{Num1: 20.4, Num2: 2.4})
	if err != nil {
		log.Fatalf("Divide RPC failed: %v", err)
	}
	fmt.Printf("Divide Result: %.2f\n", divideResult.Response)

	multiplyResult, err := client.Multiply(context.Background(), &pb.MultiplyRequest{Num1: 5.4, Num2: 2.4})
	if err != nil {
		log.Fatalf("Multiply RPC failed: %v", err)
	}
	fmt.Printf("Multiply Result: %.2f\n", multiplyResult.Response)

	arrayResult, err := client.Array(context.Background(), &pb.ArrayRequest{Nums: []float64{14.3, 2.34, 3.32}})
	if err != nil {
		log.Fatalf("Array RPC failed: %v", err)
	}
	fmt.Printf("Array Result: %.2f\n", arrayResult.Array)

	arrayDivideResult, err := client.ArrayDivide(context.Background(), &pb.ArrayDivideRequest{Nums: arrayResult.Array})
	if err != nil {
		log.Fatalf("ArrayDivide RPC failed: %v", err)
	}
	fmt.Printf("ArrayDivide Result: %.4f\n", arrayDivideResult.Array)
}
