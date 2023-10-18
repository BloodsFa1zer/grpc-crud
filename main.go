package main

import (
	"app4/config"
	"app4/database"
	"app4/handler"
	"app4/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	handler.NewServerService(grpcServer, service.NewUserService(database.NewUserDatabase()))
	cfg := config.LoadENV("config/.env")
	cfg.ParseENV()
	con, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		panic(err)
	}
	err = grpcServer.Serve(con)
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully started server")

}
