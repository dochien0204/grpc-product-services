package main

import (
	"fmt"
	"log"
	"net"
	"product_svc/pkg/config"
	"product_svc/pkg/db"
	"product_svc/pkg/pb"
	"product_svc/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed To Load Config:", err)
	}

	h := db.ConnectDB(c)

	lis, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product services on 8001")

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
