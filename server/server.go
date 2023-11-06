package server

import (
	"app4/config"
	"app4/database"
	"app4/handler"
	proto "app4/proto"
	"app4/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ServerConnection struct{}

func (sc *ServerConnection) NewServerConnection() error {
	cfg := config.LoadENV("config/.env")
	cfg.ParseENV()

	lis, err := net.Listen("tcp", "localhost:7777")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// create a server instance
	ss := handler.NewServerService(service.NewUserService(database.NewUserDatabase()))

	// attach the Ping service to the server
	proto.RegisterUserServiceServer(grpcServer, ss)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	return err
}
