package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	criticalpackage "github.com/Troelshjarne/criticalSection/critical"

	"google.golang.org/grpc"
)

var lamTime = 0

// set by flag ?
var nodeID = 0

// increments each time a client joins the cluster
var nodes = 0

// Queue of clients
var queue []int64
var queueMutex sync.Mutex

type Server struct {
	criticalpackage.UnimplementedCommunicationServer

	channel map[string][]chan *criticalpackage.Reply
}

// server reads reads node request and grant acces or put into queue
func RequestAccess(s *Server, requestStream criticalpackage.Communication_SendRequestServer) {

	go func() {
		// looking for requests in all streams
		for _, streams := range s.channel {
			// creating reply
			reply := criticalpackage.Reply{
				Access: true,
			}
			//need fix should only find channel for belonging to the requeting node
			for _, replyChan := range streams {
				// sending reply to node requesting node.
				replyChan <- &reply

				nodeId, err := requestStream.Recv()

				fmt.Println(err)
				fmt.Println(nodeId)

				queueMutex.Lock()

				// convert nodeid int64 to int ???? or change protofile to int32, random ID generator thus need to be changed
				//queue = append(queue, nodeId)
				lamTime++
				// TODO: Log "{Client} has requested access to the critical section at lamport time {Lam}."
				queueMutex.Unlock()

				fmt.Println(nodeID)
			}
		}

	}()

}

// Run in own goroutine.
// Checks intermittently, whether there is a client in queue, and serves them if possible.
func (s *Server) serveQueue() {
	if len(queue) > 0 {
		lamTime++
		// Some client is in critical section, if code is in his block.
		// TODO: Send permission for client to enter critical section.

		// TODO: Log "{client} has entered critical section at lamport time {lam}."
		// TODO: Wait here, until client gives up access.

		queueMutex.Lock()
		lamTime++
		queue = queue[1:]
		// TODO: Log "{client} has exited critical section at lamport time {lam}."
		queueMutex.Unlock()
	} else {
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {

	fmt.Println("=== Server starting up ===")
	list, err := net.Listen("tcp", ":9080")

	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	criticalpackage.RegisterCommunicationServer(grpcServer, &Server{
		channel: make(map[string][]chan *criticalpackage.Reply),
	})
	grpcServer.Serve(list)
}
