package redis

import (
	"github.com/hwangseonu/paperless.dev/internal/common"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func init() {
	redisAddr := common.GetConfig().RedisAddr
	Client = redis.NewClient(&redis.Options{Addr: redisAddr})
}
