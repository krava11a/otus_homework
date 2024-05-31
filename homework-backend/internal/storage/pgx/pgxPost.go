package pgx

import (
	"context"
	"fmt"
	"homework-backend/internal/models"
	"time"
)

// PostCreate(post models.Post) (string, error)
// PostUpdate(post_id, text string) error
// PostDelete(post_id string) error
// PostGet(post_id string) (post models.Post, err error)
// PostFeed(user_id string, offset, limit uint32) ([]models.Post, error)

func (s *Storage) PostCreate(post models.Post) (string, error) {
	const op = "storage.pgx.PostCreate"

	_, err := s.dbpool.Exec(context.Background(), `INSERT INTO posts (id_user, text, timestamp) VALUES($1,$2,$3)`, post.Id_user, post.Text, time.Now().Unix())
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	postFB, err := s.PostGetByIduserAndText(post)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return postFB.Id_post, nil
}

func (s *Storage) PostGetByIduserAndText(post models.Post) (models.Post, error) {
	const op = "storage.pgx.PostCreate"
	newPost := models.Post{}
	resRow := s.dbpool.QueryRow(context.Background(), `SELECT id_post FROM posts WHERE id_user = $1 AND text = $2 LIMIT 1`, post.Id_user, post.Text)
	resRow.Scan(&newPost.Id_post)
	if newPost.Id_post == "" {
		return newPost, fmt.Errorf("%s, %s", op, "Post doesn't found! ")
	}
	return newPost, nil
}

func (s *Storage) PostUpdate(post_id, text string) error {
	const op = "storage.pgx.PostUpdate"

	_, err := s.dbpool.Exec(context.Background(), `UPDATE posts SET text = $1, timestamp = $2 WHERE id_post = $3`, text, time.Now().Unix(), post_id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) PostDelete(post_id string) error {
	const op = "storage.pgx.PostDelete"

	_, err := s.dbpool.Exec(context.Background(), `DELETE FROM posts WHERE id_post = $1`, post_id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) PostGet(post_id string) (post models.Post, err error) {
	const op = "storage.pgx.PostGet"

	resRow := s.dbpool.QueryRow(context.Background(), `SELECT id_post,id_user,text FROM posts WHERE id_post = $1`, post_id)
	resRow.Scan(&post.Id_post, &post.Id_user, &post.Text)
	if post.Id_post == "" {
		return models.Post{}, fmt.Errorf("%s, %s", op, "Post doesn't found! ")
	}
	return
}

func (s *Storage) PostFeed(user_id string, offset, limit uint32) (models.Posts, error) {
	const op = "storage.pgx.PostFeed"

	rows, err := s.dbpool.Query(context.Background(), `SELECT id_post,id_user,text FROM posts p INNER JOIN friends f ON p.id_user=f.id_friend WHERE f.id = $1 ORDER BY p.timestamp DESC LIMIT $2 OFFSET $3`, user_id, limit, offset)
	if err != nil {
		return models.Posts{}, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	posts := make([]models.Post, 0, 100)
	var post models.Post
	for rows.Next() {

		err := rows.Scan(&post.Id_post, &post.Id_user, &post.Text)
		// err := rows.Scan(&user)
		if err != nil {
			return models.Posts{}, err
		}
		posts = append(posts, post)

	}

	return models.Posts{Posts: posts}, nil
}

// Get all readers for this user_id.
// If user_id set 0, then return all readers of DB
func (s *Storage) PostFriends(user_id string) ([]string, error) {
	const op = "storage.pgx.PostFriends"
	fIds := make([]string, 0, 100)
	if user_id == "0" {
		rows, err := s.dbpool.Query(context.Background(), `SELECT DISTINCT id FROM friends`)
		if err != nil {
			return []string{}, fmt.Errorf("%s: %w", op, err)
		}
		defer rows.Close()

		var friend_id string
		for rows.Next() {

			err := rows.Scan(&friend_id)
			// err := rows.Scan(&user)
			if err != nil {
				return []string{}, err
			}
			fIds = append(fIds, friend_id)

		}
	} else {
		rows, err := s.dbpool.Query(context.Background(), `SELECT DISTINCT id FROM friends WHERE id_friend = $1`, user_id)
		if err != nil {
			return []string{}, fmt.Errorf("%s: %w", op, err)
		}
		defer rows.Close()
		var friend_id string
		for rows.Next() {

			err := rows.Scan(&friend_id)
			// err := rows.Scan(&user)
			if err != nil {
				return []string{}, err
			}
			fIds = append(fIds, friend_id)

		}
	}

	return fIds, nil
}
