package main

import (
	"log"
	"net"

	"github.com/smoothbronco/community-station-api/location/pb"
	"github.com/smoothbronco/community-station-api/location/repository"
	"github.com/smoothbronco/community-station-api/location/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50000")

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	repository, err := repository.NewMysqlRepo()
	if err != nil {
		log.Fatalf("Failed to create repository: %v\n", err)
	}

	service := service.NewService(repository)

	server := grpc.NewServer()
	pb.RegisterLocationServiceServer(server, service)

	log.Println("Listening on 50000...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
