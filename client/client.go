package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	criticalpackage "github.com/Troelshjarne/criticalSection/critical"

	"google.golang.org/grpc"
)

//Flag to set which IP address to connect to
var tcpServer = flag.String("server", ":9080", "TCP Server")

//Random generator for unique IDs
var idGenerator, _ = rand.Int(rand.Reader, big.NewInt(10000000))
var nodeID = idGenerator.Int64()

var prevValue = int64(0)
var incValue = int64(0)

var executionCount = int64(1)

//Global variable to see if client/node has access to critical section
var nodeHasAccess = false

func main() {

	fmt.Println("Client connecting to the server!")

	//Starting TCP connection with options.
	var options []grpc.DialOption
	options = append(options, grpc.WithBlock(), grpc.WithInsecure())

	//Accessing the TCP with the flag from command line initiation
	conn, err := grpc.Dial(*tcpServer, options...)

	if err != nil {
		log.Fatalf("Failed to dial %v", err)
	}

	//When the client is closed, the connection is closed as well.
	defer conn.Close()

	//Global variable to share state
	ctx := context.Background()

	//Makes communication possible from client
	client := criticalpackage.NewCommunicationClient(conn)

	//Every 5 seconds a client should try to access the critical section
	//Requests are sent here
	for {
		time.Sleep(time.Second * 5)
		log.Println("Sending request for critical access")
		var valGen, _ = rand.Int(rand.Reader, big.NewInt(100))
		incValue = valGen.Int64() + 3*executionCount

		fmt.Printf("Prev val = %v, inc val = %v \n", prevValue, incValue)

		if prevValue < incValue {
			sendRequest(ctx, client, nodeID, incValue) //Fix ID sent with message
			prevValue = incValue
		} else {
			incValue = 0
			sendRequest(ctx, client, nodeID, incValue) //Fix ID sent with message
		}
		executionCount++
	}

}

//This function sends a request to the server, to grant access to the critical section.
//The request is simply its ID. The ID should be received by the server and put in the queue
//of clients waiting for critical section access.
func sendRequest(ctx context.Context, client criticalpackage.CommunicationClient, nodeID int64, val int64) {

	//Stream for sending requests to server
	stream, err := client.SendRequest(ctx)

	if err != nil {
		log.Printf("Failure sending request. Got this error: %v", err)
	}

	//The request is only the clients ID.
	request := criticalpackage.Request{
		NodeId: nodeID, //Should be a pointer to the id?
		Val:    incValue,
	}

	//Send the request to the server
	stream.Send(&request)

	acc, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Cant read")
	}

	if acc.Status == "Granted" {
		nodeHasAccess = true

		log.Println("Entering critical section")
		// Simulate doing stuff with critical access
		time.Sleep(5 * time.Second)
		log.Println("Exiting critical section")

		nodeHasAccess = false
	} else {
		log.Printf("Denied access to critical section. Status recieved: %s", acc.Status)
	}
}
