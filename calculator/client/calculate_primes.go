package main

import (
	"context"
	"io"
	"log"

	pb "github.com/HyungJunGoo/grpc-go/calculator/proto"
)

func doPrimesCalculate(c pb.CalculatorServiceClient) {
	log.Println("doPrimesCalculate invoked")
	req := &pb.PrimeCalculatorRequest{
		Num: 120,
	}
	stream, err := c.CalculatePrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling CalculatePrimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the strema: %v\n", err)
		}

		log.Printf("CalculatePrimes factor : %s\n", msg.String())
	}
}
