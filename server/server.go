package main

import (
	"context"
	"log"
	"net"

	"github.com/rof20004/mynewgrpctest/api"
	"google.golang.org/grpc"
)

// Server struct to call function
type Server struct{}

// CalculateDiscount show message in server
func (s *Server) CalculateDiscount(context context.Context, in *api.Request) (*api.Response, error) {
	log.Printf("Receive message %s", in.Name)
	return &api.Response{Id: in.Id, Name: in.Name, Pct: 5}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	api.RegisterDiscountServer(srv, &Server{})
	srv.Serve(listener)
}
