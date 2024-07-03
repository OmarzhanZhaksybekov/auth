package service

import (
	"github.com/ShawaDev/auth/pkg/model"
	"github.com/ShawaDev/auth/pkg/repository"
)

// service for authorization
type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(email, password string) (string, string, error)
}

type Service struct {
	Authorization
}

func NewService(repo repository.Authorization) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
