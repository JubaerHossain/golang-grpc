package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/JubaerHossain/golang-grpc/api/proto"
	"github.com/JubaerHossain/golang-grpc/internal/server/handlers"
)

func RunGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	usersHandler := handlers.NewUsersServiceHandler()

	users.RegisterUsersServiceServer(server, usersHandler)

	log.Println("Server listening on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
