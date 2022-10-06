package main

import (
	"context"
	pb "github.com/mbcarruthers/gRPCprimes/primes/proto"
	"io"
	"log"
)

func primesReceive(c pb.PrimesServiceClient) {
	log.Println("primesReceive invoked")
	req := &pb.PrimesRequest{
		PrimesArg: int64(120),
	}

	if stream, err := c.Calculate(context.Background(),
		req); err != nil {
		log.Fatalf("client-side error (calculate):\n%s\n",
			err.Error())
	} else {
		log.Printf("Factors of %d ", req.PrimesArg)
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("Error reading stream\n%s\n",
					err.Error())
			} else {
				log.Printf("%d \n", msg.PrimesResult)
			}
		}
	}
}
