package main

import (
	"context"
	"log"
	"time"

	pb "github.com/HyungJunGoo/grpc-go/calculator/proto"
)

func doAvgCalculate(c pb.CalculatorServiceClient) {
	log.Println("doAvgCalculate was invoked")

	reqs := []*pb.PrimeCalculatorRequest{
		{Num: 1},
		{Num: 2},
		{Num: 3},
		{Num: 4},
	}

	stream, err := c.CalculateAvg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling CalculateAvg %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req : %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from CalculateAvg: %v\n", err)
	}

	log.Printf("CalculateAvg: %v\n", res.Result)
}
