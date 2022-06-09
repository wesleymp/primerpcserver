package main

import (
	"log"
	"net"

	"github.com/wesleymp/primerpcserver/pb"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.PrimeServiceServer
}

func (s *Server) PrimeNumber(req *pb.PrimeNumberRequest, stream pb.PrimeService_PrimeNumberServer) error {
	N := req.PrimeNumber
	K := int32(2)
	for N > 1 {
		if N%K == 0 {
			stream.Send(&pb.PrimeNumberResponse{
				NumberResult: K,
			})
			N = N / K
		} else {
			K = K + 1
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Listen error: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterPrimeServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
