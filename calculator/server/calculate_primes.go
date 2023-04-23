package main

import (
	"log"

	pb "github.com/HyungJunGoo/grpc-go/calculator/proto"
)

func (s *Server) CalculatePrimes(in *pb.PrimeCalculatorRequest, stream pb.CalculatorService_CalculatePrimesServer) error {
	log.Printf("CalculatePrimes function was invoked with: %v\n", in)
	n := in.GetNum()
	k := 2
	for n > 1 {
		if n%int32(k) == 0 {
			n /= int32(k)
			stream.Send(
				&pb.CalculatorResponse{
					Result: int32(k),
				},
			)
		} else {
			k += 1
		}
	}
	return nil
}
