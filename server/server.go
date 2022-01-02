package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"

	criticalpackage "github.com/Troelshjarne/criticalSection/critical"

	"google.golang.org/grpc"
)

var lamTime = 0
var critVal = 0

// Queue of clients
var queue []int64
var queueMutex sync.Mutex

type Server struct {
	criticalpackage.UnimplementedCommunicationServer

	channel map[int64][]chan *criticalpackage.Reply
}

func LogSetup() {
}

func Log(text string) {
	log.Printf("Lamport Time %d %s", lamTime, text)
}

//Server reads node request and grant acces or put into queue
func (s *Server) SendRequest(requestStream criticalpackage.Communication_SendRequestServer) error {
	request, err := requestStream.Recv()
	if err != nil {
		log.Printf("Request error: %v \n", err)
	}
	nodeID := request.NodeId
	incVal := request.Val

	//What are we using this for?
	authchannel := make(chan *criticalpackage.Reply)

	queueMutex.Lock()
	queue = append(queue, nodeID)
	lamTime++

	Log(fmt.Sprintf("Client \"%v\" has requested access to the critical section and has been put in the back of the queue", nodeID))
	queueMutex.Unlock()

	go func() {
		for {
			if queue[0] == nodeID {

				// Some client is in critical section, if code is in this block.
				lamTime++

				//What are we using this for?
				authchannel <- nil

				Log(fmt.Sprintf("Client \"%v\" has entered the critical section", nodeID))

				// Simulate being in critical section, and exiting it again.
				time.Sleep(5 * time.Second)

				queueMutex.Lock()
				lamTime++

				critVal += int(incVal)

				queue = queue[1:]
				Log(fmt.Sprintf("Client \"%v\" has exited the critical section", nodeID))
				Log(fmt.Sprintf("\nThe critical value is now: %v \n", critVal))

				queueMutex.Unlock()
				break
			} else {
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()

	//What are we using this for?
	<-authchannel

	requestStream.SendAndClose(&criticalpackage.Ack{Status: "Granted"})

	return nil
}

// Run in own goroutine.
// Checks intermittently, whether there is a client in queue, and serves them if possible.
func (s *Server) serveQueue() {
	for {
		if len(queue) > 0 {

			// Some client is in critical section, if code is in this block.
			lamTime++

			id := queue[0]
			s.channel[id][0] <- nil

			Log(fmt.Sprintf("Client \"%v\" has entered the critical section", id))

			queueMutex.Lock()
			lamTime++
			queue = queue[1:]

			Log(fmt.Sprintf("Client \"%v\" has exited the critical section", id))
			queueMutex.Unlock()
		} else {
			time.Sleep(time.Millisecond * 50)
		}
	}
}

func main() {

	fmt.Println("=== Server starting up ===")
	list, err := net.Listen("tcp", ":9080")

	//Logging
	LOG_FILE := "./server.log"

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(mw)

	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}

	var options []grpc.ServerOption
	grpcServer := grpc.NewServer(options...)

	criticalpackage.RegisterCommunicationServer(grpcServer, &Server{
		channel: make(map[int64][]chan *criticalpackage.Reply),
	})
	grpcServer.Serve(list)

}
