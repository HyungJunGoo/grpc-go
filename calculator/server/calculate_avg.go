package main

import (
	"io"
	"log"

	pb "github.com/HyungJunGoo/grpc-go/calculator/proto"
)

func (s *Server) CalculateAvg(stream pb.CalculatorService_CalculateAvgServer) error {
	log.Printf("CalculateAvg was invoked")

	res := 0
	cnt := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: (float32(res) / float32(cnt)),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream :%v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		res += int(req.GetNum())
		cnt += 1
	}

}
