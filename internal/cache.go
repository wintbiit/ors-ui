package internal

import "github.com/redis/go-redis/v9"

type cache struct {
	*redis.Client
}

func bootCache() {
}
