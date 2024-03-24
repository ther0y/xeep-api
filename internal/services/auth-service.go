package services

import (
	"context"
	"github.com/ther0y/xeep-api/auther"
)

type AuthService interface {
	Login(ctx context.Context, identifier string, password string) (*auther.AuthenticationData, error)
}

func NewAuthService() AuthService {
	return newAutherGrpcService()
}
