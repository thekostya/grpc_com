package main

import (
	"context"
	"github.com/thekostya/grpc_com/lib/server"
	"google.golang.org/grpc"
	"log"
)

func main() {
	grpcClient, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect to server: %v", err)
	}
	client := server.NewTransactionsClient(grpcClient)

	transaction := &server.NewTransaction{
		Env:         server.NewTransaction_RU,
		Company:     &server.Company{Id: 35, Name: "New Test Company"},
		Transaction: &server.Transaction{Id: "New Test Transaction ID", Amount: 1000},
		User:        &server.User{Username: "thekostya"},
		Order:       &server.Order{Id: 78432},
		Calculation: &server.Calculation{Id: 234},
	}
	_, err = client.Add(context.Background(), transaction)
	if err != nil {
		log.Fatalf("Couldn't add transaction: %v", err)
	}
}
