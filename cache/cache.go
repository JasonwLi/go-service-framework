package cache

import (
	"context"
	"github.com/coherentopensource/go-service-framework/util"
	"github.com/go-redis/redis/v8"
)

type error interface {
	Error() string
}

type NotInRedisCacheError struct {
	message string
}

func (e *NotInRedisCacheError) Error() string {
	return e.message
}

type RedisConfig struct {
	Host     string
	Username string
	Password string
	DB       int
}

type Cache struct {
	redisDB *redis.Client
}

func NewRedisClient(cfg *RedisConfig, logger util.Logger) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// ping
	status := rdb.Ping(context.Background())
	if status.Err() != nil {
		logger.Fatalf("could not connect to Redis with host: %s, error: %v", cfg.Host, status.Err())
	}
	logger.Infof("connected to redis: %s", status.Val())

	return &Cache{
		rdb,
	}
}

func (r *Cache) Get(ctx context.Context, key string) (string, error) {
	strCmd := r.redisDB.Get(ctx, key)
	if strCmd.Err() != nil {
		if strCmd.Err() == redis.Nil {
			return "", &NotInRedisCacheError{message: "key does not exist in redis"}
		}
		return "", strCmd.Err()
	}
	return strCmd.Val(), nil
}

func (r *Cache) Set(ctx context.Context, key string, value interface{}) error {
	// set key without expiration
	err := r.redisDB.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
