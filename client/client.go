/* package main

import (
	"fmt"
	"log"

	criticalpackage "github.com/Troelshjarne/criticalSection/critcal"

	"google.golang.org/grpc"
)

func main () {
	fmt.Println("test")

	var options []grpc.DialOption
	options = append(options, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, options...)
	if err != nil {
		log.Fatalf("Failed to dial %v", err)
	}


}
*/