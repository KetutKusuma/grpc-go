package main

import (
	"log"

	pb "github.com/KetutKusuma/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect; %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	// log.Printf("Server started : %v", conn.())

	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob", "Mama", "Papa", "Uhuy"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
}
