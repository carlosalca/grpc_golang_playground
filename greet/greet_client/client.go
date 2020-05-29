package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"github.com/carlosalca/grpc_golang_playground/greet/greetpb"

)

func main ()  {
	fmt.Println("Hello I'm a client!")
	cc,err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	//When we're done with the connection close it
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println("Created client: %f", c)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC")
	req := &greetpb.GreetRequest {
		Greeting: &greetpb.Greeting{
			FirstName: "Carlos",
			LastName: "Alca",
		},
	}
	res,err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while callinig Greet RPC: %v",err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}