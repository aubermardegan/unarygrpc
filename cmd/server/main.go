package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	personpb "github.com/aubermardegan/unarygrpc/pb/person/v1"
)

type personServer struct {
	personpb.UnimplementedPersonServiceServer
	mu     sync.RWMutex
	people map[string]*personpb.Person
}

func (s *personServer) AddPerson(ctx context.Context, req *personpb.AddPersonRequest) (*personpb.AddPersonResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate input
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Age <= 0 || req.Age > 120 {
		return nil, status.Error(codes.InvalidArgument, "age is required")
	}

	// Generate UUID
	id := uuid.New()

	person := &personpb.Person{
		Uuid: id.String(),
		Name: req.Name,
		Age:  req.Age,
	}

	// Store the person
	s.people[id.String()] = person

	return &personpb.AddPersonResponse{Person: person}, nil
}

func (s *personServer) GetPerson(ctx context.Context, req *personpb.GetPersonRequest) (*personpb.GetPersonResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, "uuid is required")
	}

	// Retrieve the person from storage
	person, ok := s.people[req.Uuid]
	if !ok {
		return nil, status.Error(codes.NotFound, "person not found")
	}

	return &personpb.GetPersonResponse{Person: person}, nil
}

func (s *personServer) ListPeople(ctx context.Context, req *personpb.ListPeopleRequest) (*personpb.ListPeopleResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Retrieve every person from storage
	people := make([]*personpb.Person, 0, len(s.people))
	for _, p := range s.people {
		people = append(people, p)
	}

	return &personpb.ListPeopleResponse{People: people}, nil
}

func main() {
	// Create a TPC Listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC Server
	s := grpc.NewServer()

	// Initialize our personServer with an in memory database to store the people
	personServer := personServer{people: make(map[string]*personpb.Person)}

	// Register our personServer on the gRPC server
	personpb.RegisterPersonServiceServer(s, &personServer)

	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
