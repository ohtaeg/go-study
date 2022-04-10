package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	server := grpc.NewServer()
	listenErr := server.Serve(listen)
	if listenErr != nil {
		log.Fatalf("failed to serve: %s", listenErr)
	}
}
