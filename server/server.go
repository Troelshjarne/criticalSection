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
var queue []string
var queueMutex sync.Mutex

type Server struct {
	criticalpackage.UnimplementedCommunicationServer

	channel map[string][]chan *criticalpackage.Request
}

func (s *Server) RequestAccess(clientID string) {
	queueMutex.Lock()
	queue = append(queue, clientID)
	lamTime++
	// TODO: Log "{Client} has requested access to the critical section at lamport time {Lam}."
	queueMutex.Unlock()
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
		time.Sleep(time.Millisecond*50)
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
		channel: make(map[string][]chan *criticalpackage.Request),
	})
	grpcServer.Serve(list)
}
