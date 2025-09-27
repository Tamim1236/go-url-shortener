package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
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

// to save mapping between the originalURL and the generated short URL
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

// retrieve initial long URL when short URL is provided
// user calls shortlink in the URL, so retrieve long URL and then redirect
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl; %s\n", err, shortUrl))
	}
	return result
}
