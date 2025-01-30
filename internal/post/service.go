package post

import (
	"context"
	"errors"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/marco-fabian/api-crud-go/internal"
)

var ErrPostBodyEmpty = errors.New("post body is empty")
var ErrPostBodyExceedsLimit = errors.New("post body exceeds limit")
var ErrPostNotFound = errors.New("post not found")

type Service struct {
	Repository Repository
}

func (s Service) Create(post internal.Post) error {
	if post.Body == "" {
		return ErrPostBodyEmpty
	}

	if utf8.RuneCountInString(post.Body) > 140 {
		return ErrPostBodyExceedsLimit
	}

	if post.Title == "" || post.Author == "" {
		return errors.New("Os campos title e author são obrigatórios")
	}

	return s.Repository.Insert(post)
}

func (s *Service) Delete(id uuid.UUID) error {
	ctx := context.Background()

	result, err := s.Repository.Conn.Exec(ctx, "DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrPostNotFound
	}

	return nil
}

func (s *Service) List() ([]internal.Post, error) {
	ctx := context.Background()
	rows, err := s.Repository.Conn.Query(ctx, "SELECT id, username, body, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []internal.Post
	for rows.Next() {
		var post internal.Post
		err := rows.Scan(&post.ID, &post.Username, &post.Body, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Service) Update(id uuid.UUID, post internal.Post) error {
	ctx := context.Background()
	query := `UPDATE posts SET username = '` + post.Username + `', body = '` + post.Body + `' WHERE id = '` + id.String() + `'`

	_, err := s.Repository.Conn.Exec(ctx, query)
	return err
}
