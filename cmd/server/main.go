package main

import (
	"log"
	"net"

	"grpc-go-examples-bank/internal/bank/account"
	"google.golang.org/grpc"

	pb "grpc-go-examples-bank/pkg/proto/bank/account"
)

func main() {
	log.Println("Server running ...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	pb.RegisterDepositServiceServer(server, &account.DepositServer{})

	log.Fatalln(server.Serve(listener))
}
