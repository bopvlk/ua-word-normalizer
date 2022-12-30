package scraper

import (
	"log"

	"github.com/go-redis/redis"
)

// var badWords []string

func normalizeWord(rdb *redis.Client, word string) string {
	goodWord, dberr := GetFromRedis(rdb, word)

	if dberr == nil {
		log.Println("this word exist in database")
		return goodWord

	}

	goodWord, scrapErr := R2UScparer(word)

	if dberr == redis.Nil {
		SetToRedis(rdb, word, goodWord)
	} else {
		log.Println("redis error:", dberr)
		if scrapErr != nil {
			log.Println("scraper error:", scrapErr)
			return word
		}
	}

	return goodWord
}

func MakingGoodWords(badWords []string) {
	rdb := NewRedisClient()
	for i := 0; i < len(badWords); i++ {
		badWords[i] = normalizeWord(rdb, badWords[i])
	}
}
