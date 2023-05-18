package trantor

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type RedisWrapper struct {
	redisClient *redis.Client
	ctx         context.Context
}

// Value Redis Query
func ValidateQuery(key string, value string) bool {
	if key == "" || value == "" {
		return false
	}
	return true
}

// Make Custom Redis Wrapper
func NewRedisWrapper() (*RedisWrapper, error) {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})

	_, err := redis.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisWrapper{
		redisClient: redis,
		ctx:         context.Background(),
	}, nil
}

// Close Redis Connection
func (r *RedisWrapper) Close() error {
	r.redisClient.Close()
	return nil
}

// Read Redis Value
func (r *RedisWrapper) Read(key string) (string, error) {

	if key == "" {
		return "", errors.New("key is empty")
	}

	val, err := r.redisClient.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

// Write Redis Value
func (r *RedisWrapper) Write(key string, value string) error {

	if key == "" {
		return errors.New("key is empty")
	}

	err := r.redisClient.Set(r.ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// Delete Redis Value
func (r *RedisWrapper) Delete(key string) error {

	if key == "" {
		return errors.New("key is empty")
	}

	err := r.redisClient.Del(r.ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
