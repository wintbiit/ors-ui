package internal

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type cache struct {
	*redis.Client
}

var Cache *cache

func bootCache() {
	Cache = &cache{
		Client: redis.NewClient(&redis.Options{
			Addr:     Config.Redis.Address,
			Password: Config.Redis.Password,
			DB:       3,
		}),
	}

	_, err := Cache.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal().Msgf("redis ping failed: %v", err)
	}

	log.Info().Msg("Redis connected")
}
