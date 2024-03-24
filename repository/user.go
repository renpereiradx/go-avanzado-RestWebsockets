package repository

import (
	"context"

	"github.com/renpereiradx/go-avanzado-RestWebsocket/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	Close() error
}

var implementation UserRepository

func SetUserRepository(user UserRepository) {
	implementation = user
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserByID(ctx, id)
}
