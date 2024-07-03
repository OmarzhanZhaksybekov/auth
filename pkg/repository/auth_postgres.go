package repository

import (
	"github.com/ShawaDev/auth/pkg/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := "INSERT INTO users (email, phone, password, role) VALUES ($1, $2, $3, $4) RETURNING id"

	if user.Role == "" {
		user.Role = "user"
	}

	row := r.db.QueryRow(query, user.Email, user.Phone, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE email = $1 AND password = $2"
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		return model.User{}, err
	}

	return user, err
}
