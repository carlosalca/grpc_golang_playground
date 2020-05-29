package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"github.com/carlosalca/grpc_golang_playground/calculator/calculatorpb"

)

func main ()  {
	fmt.Println("Calculator client!")
	cc,err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	//When we're done with the connection close it
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	fmt.Println("Created client: %f", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC")
	req := &calculatorpb.SumRequest {
		FirstNumber: 5,
		SecondNumber: 40,
	}
	res,err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while callinig Greet RPC: %v",err)
	}
	log.Printf("Response from Greet: %v", res.SumResult)
}