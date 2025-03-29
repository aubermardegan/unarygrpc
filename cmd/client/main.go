package main

import (
	"context"
	"log"
	"time"

	personpb "github.com/aubermardegan/unarygrpc/pb/person/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Creates a new client connection on localhost:50051
	clientConn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer clientConn.Close()

	// Creates a PersonServiceClient using the client connection
	client := personpb.NewPersonServiceClient(clientConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Add a person
	addResp, err := client.AddPerson(ctx, &personpb.AddPersonRequest{
		Name: "Auber",
		Age:  35,
	})
	if err != nil {
		log.Fatalf("AddPerson failed: %v", err)
	}
	log.Printf("Added person: %v", addResp.Person)

	// Get the person
	getResp, err := client.GetPerson(ctx, &personpb.GetPersonRequest{
		Uuid: addResp.Person.Uuid,
	})
	if err != nil {
		log.Fatalf("GetPerson failed: %v", err)
	}
	log.Printf("Got person: %v", getResp.Person)

	// Add Second Person
	_, err = client.AddPerson(ctx, &personpb.AddPersonRequest{
		Name: "João",
		Age:  20,
	})
	if err != nil {
		log.Fatalf("AddPerson failed: %v", err)
	}

	// List all people
	listResp, err := client.ListPeople(ctx, &personpb.ListPeopleRequest{})
	if err != nil {
		log.Fatalf("ListPeople failed: %v", err)
	}
	log.Printf("People: %v", listResp.People)
}
