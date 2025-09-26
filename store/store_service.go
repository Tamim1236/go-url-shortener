package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

// initialize store service anad return store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Reddis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}
