package main

import (
	"awesomeProject/grpcServer"
	api "awesomeProject/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	server := &grpcServer.GRPCServer{}
	api.RegisterReverserServer(s, server)

	t, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(t); err != nil {
		log.Fatal(err)
	}
}
