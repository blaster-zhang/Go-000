//+build wireinject
package main

import (
	"Go000/Week04/internal/biz"
	"Go000/Week04/internal/data"
	"Go000/Week04/internal/service"
	"context"
	"github.com/google/wire"
)

func InitUserService(ctx context.Context, repo *data.UserRepo) *service.UserService {
	wire.Bind(service.NewUserService, biz.NewUserUseCase)
	return &service.UserService{}
}
