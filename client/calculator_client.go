package main

import (
	"context"
	"log"
	"time"

	pbCalculator "github.com/andrewlawrence80/grpc-demo/proto/calculator"
	"google.golang.org/grpc"
)

func Calculate() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pbCalculator.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Example for Add
	addResp, err := client.Add(ctx, &pbCalculator.CalcRequest{Num1: 1, Num2: 2})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Add Result: %f", addResp.GetResult())

	// Example for Subtract
	subtractResp, err := client.Subtract(ctx, &pbCalculator.CalcRequest{Num1: 5, Num2: 3})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	log.Printf("Subtract Result: %f", subtractResp.GetResult())

	// Example for Multiply
	multiplyResp, err := client.Multiply(ctx, &pbCalculator.CalcRequest{Num1: 3, Num2: 4})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}
	log.Printf("Multiply Result: %f", multiplyResp.GetResult())

	// Example for Divide
	divideResp, err := client.Divide(ctx, &pbCalculator.CalcRequest{Num1: 10, Num2: 2})
	if err != nil {
		log.Fatalf("could not divide: %v", err)
	}
	log.Printf("Divide Result: %f", divideResp.GetResult())
}
