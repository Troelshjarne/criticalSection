package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"

	criticalpackage "github.com/Troelshjarne/criticalSection/critical"

	"google.golang.org/grpc"
)

var tcpServer = flag.String("server", ":5050", "TCP Server")

var idGenerator, _ = rand.Int(rand.Reader, big.NewInt(10000000))
var nodeID = idGenerator.Int64()

var nodeHasAccess = false

func main() {
	fmt.Println("test")

	var options []grpc.DialOption
	options = append(options, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, options...)
	if err != nil {
		log.Fatalf("Failed to dial %v", err)
	}

	defer conn.Close()

	ctx := context.Background()
	client := criticalpackage.NewCommunicationClient(conn)

	//go joinCluster

}

func serverReply(ctx context.Context, client criticalpackage.CommunicationClient) {

	reply := criticalpackage.Reply{Access: *&nodeHasAccess}

	stream, err := client.ServerReply(ctx, &reply)
	if err != nil {
		log.Fatalf("Client reply connection error! Throws %v", err)
	}

	waitChannel := make(chan struct{})

}

func sendRequest(ctx context.Context, client criticalpackage.CommunicationClient) {

	stream, err := client.SendRequest(ctx)

	if err != nil {
		log.Printf("Failure sending request. Got this error: %v", err)
	}

	rq := criticalpackage.Request{
		NodeId: nodeID,
	}
	stream.Send(&rq)

	ack, err := stream.CloseAndRecv()
	fmt.Println("Sent ID to server: %v \n", ack)

}
