### Fizzbuzz in golang, grpc and grpc-gateway

This is an implementation of fizzbuzz using protobuf for code generation.

It uses grpc and grpc-gateway for REST.

### How to use

#### Build

go build -o bin/fizzbuzz cmd/main.go

#### Run

bin/fizzbuzz -http-host <HTTP_ADDRESS>

go run ./cmd/main.go -http-host <HTTP_ADDRESS>

### there are 3 endpoints:

#### Swagger for documentation

```/swagger```

#### User statistics

```/stats ```

#### fizzbuzz

```/fizzbuzz```

