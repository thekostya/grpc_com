package main

import (
	"context"
	"github.com/thekostya/grpc_com/lib/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

type TransactionServer struct{}

func (ts *TransactionServer) Add(ctx context.Context, tr *server.NewTransaction) (*server.Empty, error) {
	log.Printf("Add transaction for company ID: %d", tr.Company.Id)
	return &server.Empty{}, nil
}

func main() {
	srv := grpc.NewServer()
	var transSrv *TransactionServer
	server.RegisterTransactionsServer(srv, transSrv)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Couldn't start server: %v", err)
	}
	log.Fatal(srv.Serve(l))
}
