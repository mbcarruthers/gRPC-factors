package main

import (
	pb "github.com/mbcarruthers/gRPCprimes/primes/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	addr = "0.0.0.0:5051"
	tcp  = "tcp"
)

type Server struct {
	pb.PrimesServiceServer
}

func main() {
	lis, err := net.Listen(tcp, addr)
	if err != nil {
		log.Fatalf("error listening at %s:%s\n %s \n",
			tcp, addr, err.Error())
	}
	log.Printf("[Server] Listening at %s:%s",
		tcp, addr)
	srv := grpc.NewServer()
	pb.RegisterPrimesServiceServer(srv, &Server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("grpc listen error \n %s \n",
			err.Error())
	}
}