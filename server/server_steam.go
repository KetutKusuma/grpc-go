package main

import (
	"log"
	"time"

	pb "github.com/KetutKusuma/grpc-go/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names : %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		//delayed
		time.Sleep(2 * time.Second)

	}

	return nil
}
