package services

import (
	"context"
	"github.com/ther0y/xeep-api/auther"
	"github.com/ther0y/xeep-api/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type autherGrpc struct {
	cc grpc.ClientConnInterface
}

type autherGrpcService struct {
	client auther.AutherClient
}

func newAutherGrpcService() AuthService {
	autherUrl, err := utils.GetEnv("AUTHER_URL")
	if err != nil {
		log.Fatal("AUTHER_URL is not set")
	}

	conn, err := grpc.Dial(autherUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Auther did not connect: %v", err)
	}

	client := auther.NewAutherClient(conn)

	return &autherGrpcService{
		client,
	}
}

func (s *autherGrpcService) Login(ctx context.Context, identifier string, password string) (*auther.AuthenticationData, error) {
	res, err := s.client.Login(ctx, &auther.LoginRequest{
		Identifier: identifier,
		Password:   password,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
