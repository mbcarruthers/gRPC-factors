package main

import (
	pb "github.com/mbcarruthers/gRPCprimes/primes/proto"
	"log"
)

func (s *Server) Calculate(in *pb.PrimesRequest,
	stream pb.PrimesService_CalculateServer) error {
	log.Printf("Calculat called from %v \n", in)

	value := in.PrimesArg
	var k int64 = 2

	for value > 1 {
		if value%k == 0 {
			_ = stream.Send(&pb.PrimesResponse{
				PrimesResult: k,
			})
			value = value / k
		} else {
			k++
		}
	}
	return nil
}
