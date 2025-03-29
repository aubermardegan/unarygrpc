# Unary gRPC Example in Go

This repository contains an example of a unary gRPC service implementation in Go!

## Overview

The example implements a simple `PersonService` with three methods:

- `AddPerson`: Creates a new person record with server-generated UUID
- `GetPerson`: Retrieves a person by UUID
- `ListPeople`: Returns all stored person records

## Prerequisites

### 1. Install Protocol Buffer Compiler (protoc)

Download and install the Protocol Buffer compiler for your platform:

- [Download protoc](https://github.com/protocolbuffers/protobuf/releases)

### 2. Install Go Dependencies

Run the following commands to install required Go dependencies:

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.71.0
go get google.golang.org/grpc
go get github.com/google/uuid
```

## Generating Go Code

To generate the Go code from the Protocol Buffer definitions, run:

```shell
protoc --go_out=. --go-grpc_out=. proto/v1/person.proto
```

This will generate the following files in the pb/person/v1/ directory:

- person.pb.go: Protocol Buffer message definitions
- person_grpc.pb.go: gRPC server and client interfaces

## Project Structure

```
.
├── cmd/
│   ├── client/         # Client implementation
│   │   └── main.go
│   └── server/         # Server implementation
│       └── main.go
├── proto/
│   └── v1/             # Protocol Buffer definitions
│       └── person.proto
└── pb/
    └── person/
        └── v1/         # Generated code
            ├── person.pb.go
            └── person_grpc.pb.go
```

## Running the Example

1. Start the gRPC server:

```shell
go run cmd/server/main.go
```

2. In a separate terminal, run the client:

```shell
go run cmd/client/main.go
```
