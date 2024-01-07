package client

import (
	"app4/config"
	proto "app4/proto"
	_ "context"
	"google.golang.org/grpc"
	"log"
)

type ClientService struct {
	clientConn *proto.UserServiceClient
}

func NewClientService() *ClientService {
	cfg := config.LoadENV("config/.env")
	cfg.ParseENV()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewUserServiceClient(conn)

	return &ClientService{clientConn: &c}
}
