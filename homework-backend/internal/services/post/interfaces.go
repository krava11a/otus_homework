package post

import (
	"homework-backend/internal/models"
	"time"
)

type PostCreater interface {
	PostCreate(post models.Post) (string, error)

	PostUpdate(post_id, text string) error
	PostDelete(post_id string) error

	FriendSet(user_id, friend_id string) error
	FriendDelete(user_id, friend_id string) error
}

type PostProvider interface {
	PostGet(post_id string) (post models.Post, err error)
	PostFeed(user_id string, offset, limit uint32) (models.Posts, error)

	PostFriends(user_id string) ([]string, error)
}

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, value interface{}) error
}

type RQueuePublisher interface {
	PublishTo(name, message string) error
}
