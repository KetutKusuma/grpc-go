package main

import (
	"context"
	"io"
	"log"

	pb "github.com/KetutKusuma/grpc-go/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names : %v", err)
	}

	for {
		message, err := stream.Recv()
		// ini untuk mengecek masih ada data atau tidak
		// ((EOF is the error returned by Read when no more input is available))
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}

		log.Println(message)
	}

	log.Printf("Streaming finished")
}
