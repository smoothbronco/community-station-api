package main

import (
	"log"
	"net"
	"os"

	"github.com/smoothbronco/community-station-api/location/pb"
	"github.com/smoothbronco/community-station-api/location/repository"
	"github.com/smoothbronco/community-station-api/location/service"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	lis, err := net.Listen("tcp", ":50000")

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	dsn := os.Getenv("COMMUNITY_STATION_DEV_DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}
	db.AutoMigrate(&pb.Location{})
	db.Logger = db.Logger.LogMode(logger.Info)

	repository, err := repository.NewMysqlRepo(db)
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
