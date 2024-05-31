package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	client *redis.Client
}

func New(url string) (*DB, error) {
	const op = "storage.redis.New"

	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("%s Unable to create connection to Redis: %v\n", op, err)
	}
	client := redis.NewClient(opt)
	return &DB{client: client}, nil
}

func (d *DB) Set(key string, value interface{}, expiration time.Duration) error {
	const op = "storage.redis.Set"
	err := d.client.Set(context.Background(), key, value, expiration)
	if err != nil {
		return fmt.Errorf("%s Unable to set data to Redis: %v\n", op, err)
	}
	return nil
}

func (d *DB) Get(key string, value interface{}) error {
	const op = "storage.redis.Get"
	rn := d.client.Get(context.Background(), key)
	err := d.client.Get(context.Background(), key).Scan(value)
	if err != nil {
		return fmt.Errorf("%s Unable to get data from Redis: %v\n", op, err, rn)
	}
	return nil
	// return fmt.Errorf("%s Unable to get data from Redis: %v\n", op, p)
}
