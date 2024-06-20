package service

import (
	"github.com/ShawaDev/auth/pkg/model"
	"github.com/ShawaDev/auth/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repo repository.Authorization) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
