package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	criticalpackage "github.com/Troelshjarne/criticalSection/critical"

	"google.golang.org/grpc"
)

var lamTime = 0

// set by flag ?
var nodeID = 0

// increments each time a client joins the cluster
var nodes = 0

type Server struct {
	criticalpackage.UnimplementedCommunicationServer

	channel map[string][]chan *criticalpackage.Request
}

func main() {

	fmt.Println("=== Node starting up ===")
	list, err := net.Listen("tcp", ":9080")

	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	criticalpackage.RegisterCommunicationServer(grpcServer, &Server{
		channel: make(map[string][]chan *criticalpackage.Request),
	})
	grpcServer.Serve(list)
}

func sendRequest(ctx context.Context, client criticalpackage.CommunicationClient) {

	channel := criticalpackage.Channel{NodeId: "2"}

	stream, err := client.SendRequest(ctx, &channel)

	if err != nil {
		log.Fatalf("Client join channel error! Throws %v", err)
	}

	go func() {

		for {

			in, err := stream.RecvMsg()
			if err == io.EOF {
				close(waitChannel)
				return
			}
			if err != nil {
				log.Fatalf("Failed to recieve message from channel joining. \n Got error: %v", err)
			}

			if *senderName != in.ParticipantID {

				log.Printf("Time: (%v) Message: (%v) -> %v \n", in.LamTime, in.ParticipantID, in.Message)
			}

		}
	}()
}
