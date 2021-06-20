package main

import (
	"context"
	"log"
	"time"

	"github.com/bytejedi/grpc_example/internal/bank/account"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client running ...")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	response, err := account.
		NewDepositClient(conn, time.Second).
		Deposit(context.Background(), 1990.01)

	log.Println(response)
	log.Println(err)
}
