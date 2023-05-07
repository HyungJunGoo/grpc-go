package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/HyungJunGoo/grpc-go/calculator/proto"
)

func doMaxCalculate(c pb.CalculatorServiceClient) {
	log.Println("doMaxCalculate was invoked")

	stream, err := c.CalculateMax(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream %v\n", err)
	}

	reqs := []*pb.PrimeCalculatorRequest{
		{Num: 1},
		{Num: 5},
		{Num: 3},
		{Num: 6},
		{Num: 2},
		{Num: 20},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
			}
			log.Printf("Recieved: %v\n", res.String())
		}
		close(waitc)
	}()
	<-waitc
}
