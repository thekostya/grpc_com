package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/thekostya/grpc_com/lib/server"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	gateway := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := server.RegisterTransactionsHandlerFromEndpoint(context.Background(), gateway, "localhost:8888", opts)

	if err != nil {
		glog.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", checkAuth(gateway))

	glog.Fatal(http.ListenAndServe(":8080", mux))
}

func checkAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		...
		ctx = context.WithValue(ctx, "UserID", userID)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}