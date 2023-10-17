package client

import (
	"fmt"
	"github.com/kwanok/minsim-consumer/config"
	"google.golang.org/grpc"
	"log"
)

type GrpcClient struct {
	*grpc.ClientConn
}

func NewGrpcClient(c *config.Config) *GrpcClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.Api.Host, c.Api.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("grpc client connected")

	return &GrpcClient{
		ClientConn: conn,
	}
}
