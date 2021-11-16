package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"time"

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

	go serverReply(ctx, client)

	for i := 0; i < 100; i++ {
		fmt.Println("im alive")
		time.Sleep(time.Second * 5)
		go sendRequest(ctx, client, nodeID) //Fix ID sent with message

	}

}

func serverReply(ctx context.Context, client criticalpackage.CommunicationClient) {

	//We need to use that the boolean is send with this call
	reply := criticalpackage.Reply{Access: nodeHasAccess}

	stream, err := client.ServerReply(ctx, &reply)
	if err != nil {
		log.Fatalf("Client reply connection error! Throws %v", err)
	}

	waitChannel := make(chan struct{}) //Check if needed

	go func() {

		for {

			incomingReply, err := stream.Recv()
			if err == io.EOF {
				close(waitChannel)
				return
			}
			if err != nil {
				log.Fatalf("Failed to recieve message from sent reply. Got error: %v \n", err)
			}

			logThis(incomingReply.NodeId)

		}

	}()

	<-waitChannel

}

func sendRequest(ctx context.Context, client criticalpackage.CommunicationClient, nodeID int64) {

	stream, err := client.SendRequest(ctx)

	if err != nil {
		log.Printf("Failure sending request. Got this error: %v", err)
	}

	rq := criticalpackage.Request{
		NodeId: nodeID, //Should be a pointer to the id
	}
	stream.Send(&rq)

	ack, err := stream.CloseAndRecv()
	fmt.Println("Sent ID to server. Acknowledge = %v \n", ack)

}

func logThis(reply int64) {
	if nodeHasAccess {
		log.Println("Access = %v", reply)
	}
}
