package main

import (
	"context"
	"github.com/artback/grpc_http_fizzbuzz/pkg/fizzbuzzsrv"
	"github.com/artback/grpc_http_fizzbuzz/pkg/memory"
	"github.com/artback/grpc_http_fizzbuzz/proto/v1/fizzbuzz"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	err := fizzbuzz.RegisterFizzBuzzServiceHandlerServer(ctx, mux, &fizzbuzzsrv.Service{Statistics: memory.NewMemoryStatistics()})
	if err != nil {
		log.Fatal(err)
	}
	r := http.NewServeMux()
	r.Handle("/", mux)
	log.Println("listening to port *:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
