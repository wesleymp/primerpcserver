package main

import (
	"context"
	"io"
	"log"

	"github.com/wesleymp/primerpcserver/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

var number int32 = 120

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection error: %v\n", err)
	}
	c := pb.NewPrimeServiceClient(conn)
	req := &pb.PrimeNumberRequest{
		PrimeNumber: number,
	}
	stream, err := c.PrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("Request error: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Stream error: %v\n", err)
		}
		log.Printf("Number: %d\n", res.NumberResult)
	}
}
