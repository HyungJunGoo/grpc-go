package main

import (
	"io"
	"log"

	pb "github.com/HyungJunGoo/grpc-go/calculator/proto"
)

func (s *Server) CalculateMax(stream pb.CalculatorService_CalculateMaxServer) error {
	log.Println("CalculateMax was invoked")

	maxInt := -1 << 31

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream : %v\n", err)
		}
		if maxInt < int(req.Num) {
			maxInt = int(req.Num)
		}
		res := maxInt
		err = stream.Send(&pb.CalculatorResponse{
			Result: int32(res),
		})

		if err != nil {
			log.Fatalf("Error while sending data to client %v\n", err)
		}
	}
}
