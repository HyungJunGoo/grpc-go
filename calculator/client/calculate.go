package main

import (
	"context"
	"log"

	pb "github.com/HyungJunGoo/grpc-go/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invoked")
	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		Val1: 10,
		Val2: 3,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}
	log.Printf("Calculate result : %v\n", res.Result)
}
