package main

import (
	pb "GRPC/proto"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := os.Getenv("SERVER_ADDR")
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Yash", "Alice", "Bob"},
	}

	callSayHelloServerStream(client, names)
	callSayHello(client)
	callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
