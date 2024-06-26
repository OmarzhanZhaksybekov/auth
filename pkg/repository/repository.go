package repository

import (
	"github.com/ShawaDev/auth/pkg/model"
	"github.com/jmoiron/sqlx"
)

// functions for authorization in DB
type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(email, password string) (model.User, error)
}

// repository for interaction with db
type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
