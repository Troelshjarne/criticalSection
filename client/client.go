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

	//TO DO: Set up this function correctly
	go serverReply(client, &criticalpackage.Request{NodeId: nodeID})

	//Every x seconds (at a random interval) a client should try to access the critical section
	//Requests are sent here
	for {
		fmt.Println("im alive")
		time.Sleep(time.Second * 5)
		go sendRequest(ctx, client, nodeID) //Fix ID sent with message
	}

}

//UPDATE: This function does not stream anymore, but simply responds to messages sent.
//The big block of commented out text/code below can be deleted if we do not need streaming
//of this function anymore.
func serverReply(client criticalpackage.CommunicationClient, nodeID *criticalpackage.Request) {

	//This code should be enough, but the serverside functionality is not built in yet.
	log.Printf("Getting access status for ID: %v", nodeID)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	access, err := client.ServerReply(ctx, nodeID)
	if err != nil {
		log.Fatalf("Something bad happened with client %v getting error: %v", nodeID, err)
	}
	log.Printf("Node with ID: %v, has access status = %v \n", nodeID, access)

}

//
//
//
//
//This function should read from a stream, waiting for the server to reply with access granted or denied

//If the client receives a boolean that is false, it is not granted access
//If the client receives a boolean that is true, it is granted access

//The current problem with this function and how it is set up, is that it sends the same
//message to all clients at once. It is a stream that all clients read from. Therefore it
//is not possible with the current setup to send access to a specific client.
//We need this function to send/give access to a specific client.

/* 	stream, err := client.ServerReply(ctx, &reply)
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

}  */

//This function sends a request to the server, to grant access to the critical section.
//The request is simply its ID. The ID should be received by the server and put in the queue
//of clients waiting for critical section access.
func sendRequest(ctx context.Context, client criticalpackage.CommunicationClient, nodeID int64) {

	//Stream for listening to client requests
	stream, err := client.SendRequest(ctx)

	if err != nil {
		log.Printf("Failure sending request. Got this error: %v", err)
	}

	//The request is only the clients ID.
	rq := criticalpackage.Request{
		NodeId: nodeID, //Should be a pointer to the id?
	}
	//Send the request to the server
	stream.Send(&rq)

	//When a message is sent to the server, a client recieves an acknowledgement.
	ack, err := stream.CloseAndRecv()
	fmt.Printf("Sent ID to server. Acknowledge = %v \n", ack)

}

//For testing. Can be deleted
func logThis(reply int64) {
	if nodeHasAccess {
		log.Println("Access = %v", reply)
	}
}
