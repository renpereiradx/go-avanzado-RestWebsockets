package repository

import (
	"context"

	"github.com/renpereiradx/go-avanzado-RestWebsocket/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, posts *models.Posts) error
	GetPostByID(ctx context.Context, id string) (*models.Posts, error)
	UpdatePost(ctx context.Context, post *models.Posts, userID string) error
	DeletePost(ctx context.Context, id string, userID string) error
	ListPosts(ctx context.Context, page uint64) ([]*models.Posts, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserByID(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, posts *models.Posts) error {
	return implementation.InsertPost(ctx, posts)
}

func GetPostByID(ctx context.Context, id string) (*models.Posts, error) {
	return implementation.GetPostByID(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Posts, userID string) error {
	return implementation.UpdatePost(ctx, post, userID)
}

func DeletePost(ctx context.Context, id string, userID string) error {
	return implementation.DeletePost(ctx, id, userID)
}

func ListPosts(ctx context.Context, page uint64) ([]*models.Posts, error) {
	return implementation.ListPosts(ctx, page)
}
