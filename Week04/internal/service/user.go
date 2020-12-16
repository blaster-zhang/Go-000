package service

import (
	"Go000/Week04/internal/biz"
	"context"
)

type RequestInfo struct {
	Name         string
	Introduction string
}

type UserService struct {
	UserUseCase biz.UserUseCase
}

func NewUserService(uuc biz.UserUseCase) UserService {
	return UserService{UserUseCase: uuc}
}

func (u *UserService) Create(ctx context.Context, req RequestInfo) error {
	user := new(biz.User)
	user.Name = req.Name
	user.Introduction = req.Introduction

	return u.UserUseCase.Add(user)
}
