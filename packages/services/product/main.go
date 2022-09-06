package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ThatGuyKev/ecommerce-services/product/db"
	"github.com/ThatGuyKev/ecommerce-services/product/handlers"
	pb "github.com/ThatGuyKev/ecommerce-services/product/pb"
	"google.golang.org/grpc"
)

type Config struct {
	Port  string
	DBUrl string
}

func main() {
	config := &Config{
		Port:  "0.0.0.0:9101",
		DBUrl: "postgresql://postgres:password@0.0.0.0:9102/db",
	}

	h := db.Init(config.DBUrl)
	lis, err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatalln("Server failed to establish TCP connection", err)
	}

	fmt.Println("Product Services listening on: ", config.Port)

	s := handlers.Server{
		H: h,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to establish grpc connection: ", err)
	}
}
