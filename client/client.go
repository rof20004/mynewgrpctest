package main

import (
	"encoding/json"
	"log"

	"github.com/rof20004/mynewgrpctest/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewDiscountClient(conn)
	response, err := c.CalculateDiscount(context.Background(), &api.Request{Id: 1, Name: "TV de 42 polegadas da Sony"})
	if err != nil {
		log.Fatalf("Error when calling CalculateDiscount: %s", err)
	}

	b, _ := json.Marshal(response)

	log.Printf("Response from server: %s", string(b))
}
