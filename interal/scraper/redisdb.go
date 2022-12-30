package scraper

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func NewRedisClient() *redis.Client {
	godotenv.Load()
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: os.Getenv("REDIS"),
		DB:       0, // use default DB
	})
}

func SetToRedis(rdb *redis.Client, bedWord string, goodWord string) error {
	err := rdb.Set(bedWord, goodWord, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetFromRedis(rdb *redis.Client, bedWord string) (string, error) {
	return rdb.Get(bedWord).Result()

}
