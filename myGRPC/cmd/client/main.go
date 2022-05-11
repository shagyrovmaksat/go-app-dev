package main

import (
	api "awesomeProject/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("Error")
	}
	message := flag.Arg(0)

	connection, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewReverserClient(connection)
	response, err := client.Reverse(context.Background(), &api.ReverseRequest{Message: message})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.GetResult())
}
