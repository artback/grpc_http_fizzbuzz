package main

import (
	"context"
	"flag"
	"github.com/artback/grpc_http_fizzbuzz/pkg/fizzbuzzsrv"
	"github.com/artback/grpc_http_fizzbuzz/pkg/memory"
	"github.com/artback/grpc_http_fizzbuzz/proto/v1/fizzbuzz"
	"github.com/flowchartsman/swaggerui"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kuangchanglang/graceful"
	"log"
	"net/http"
	"os"
)

const (
	HttpDefaultAddr = ":8080"
)

var (
	httpHost = flag.String("http-host", HttpDefaultAddr, "HTTP address")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	flag.Parse()

	swaggerFile, err := os.ReadFile("proto/v1/fizzbuzz/fizzbuzz.swagger.json")
	if err != nil {
		log.Fatal(err)
	}

	mux := runtime.NewServeMux()
	err = fizzbuzz.RegisterFizzBuzzServiceHandlerServer(ctx, mux, &fizzbuzzsrv.Service{Statistics: memory.NewMemoryStatistics()})
	if err != nil {
		log.Fatal(err)
	}
	r := http.NewServeMux()
	r.Handle("/", mux)
	r.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(swaggerFile)))
	server := graceful.NewServer()
	server.Register(*httpHost, r)
	log.Println("listening on address " + *httpHost)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
