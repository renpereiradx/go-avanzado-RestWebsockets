package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/renpereiradx/go-avanzado-RestWebsocket/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (p *PostgresRepository) Close() error {
	return p.db.Close()
}

func (p *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

func (p *PostgresRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println("error closing rows", err)
		}
	}()
	var user models.User
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}
func (p *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println("error closing rows", err)
		}
	}()
	var user models.User
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *PostgresRepository) InsertPost(ctx context.Context, posts *models.Posts) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO posts (id, post_content, user_id) VALUES ($1, $2, $3)", posts.ID, posts.PostContent, posts.UserID)
	return err
}

func (p *PostgresRepository) GetPostByID(ctx context.Context, id string) (*models.Posts, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, post_content, created_at, user_id FROM posts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println("error closing rows", err)
		}
	}()
	var post = models.Posts{}
	for rows.Next() {
		if err = rows.Scan(&post.ID, &post.PostContent, &post.CreatedAt, &post.UserID); err == nil {
			return &post, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *PostgresRepository) UpdatePost(ctx context.Context, post *models.Posts, userID string) error {
	_, err := p.db.ExecContext(ctx, "UPDATE posts SET post_content = $1 WHERE id = $2 AND user_id = $3", post.PostContent, post.ID, userID)
	return err
}

func (p *PostgresRepository) DeletePost(ctx context.Context, id string, userID string) error {
	_, err := p.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 AND user_id = $2", id, userID)
	return err
}
