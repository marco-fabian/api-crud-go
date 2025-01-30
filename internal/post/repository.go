package post

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/marco-fabian/api-crud-go/internal"
)

type Repository struct {
	Conn *pgxpool.Pool
}

func (r *Repository) Insert(post internal.Post) error {
	_, err := r.Conn.Exec(
		context.Background(),
		"INSERT INTO posts (username, body, title, author) VALUES ($1, $2, $3, $4)",
		post.Username,
		post.Body,
		post.Title,
		post.Author,
	)

	return err
}

func (r *Repository) Delete(id uuid.UUID) error {
	result, err := r.Conn.Exec(context.Background(), "DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("Erro ao deletar post: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrPostNotFound
	}

	return nil
}

func (r *Repository) List() ([]internal.Post, error) {
	rows, err := r.Conn.Query(context.Background(), "SELECT id, username, body, title, author, created_at FROM posts")
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar posts: %w", err)
	}
	defer rows.Close()

	var posts []internal.Post
	for rows.Next() {
		var post internal.Post
		err := rows.Scan(&post.ID, &post.Username, &post.Body, &post.Title, &post.Author, &post.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("Erro ao ler post: %w", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Erro durante a iteração dos posts: %w", err)
	}

	return posts, nil
}

func (r *Repository) Update(id uuid.UUID, post internal.Post) error {
	query := "UPDATE posts SET "
	args := []interface{}{}
	counter := 1

	if post.Username != "" {
		query += fmt.Sprintf("username = $%d, ", counter)
		args = append(args, post.Username)
		counter++
	}

	if post.Body != "" {
		query += fmt.Sprintf("body = $%d, ", counter)
		args = append(args, post.Body)
		counter++
	}

	if post.Title != "" {
		query += fmt.Sprintf("title = $%d, ", counter)
		args = append(args, post.Title)
		counter++
	}

	if post.Author != "" {
		query += fmt.Sprintf("author = $%d, ", counter)
		args = append(args, post.Author)
		counter++
	}

	if len(args) == 0 {
		return errors.New("Campo não encontrado para atualização")
	}

	query = query[:len(query)-2] + " WHERE id = $" + fmt.Sprint(counter)
	args = append(args, id)

	_, err := r.Conn.Exec(context.Background(), query, args...)
	return err
}
